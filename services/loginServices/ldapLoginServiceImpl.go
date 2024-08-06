package loginservices

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lenna-ai/bni-iproc/helpers/jwtHelpers/decrypt"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
	usermodel "github.com/lenna-ai/bni-iproc/models/userModel"
	"github.com/shirou/gopsutil/host"
)



func (ldapLoginServiceImpl *LdapLoginServiceImpl) AuthUsingLDAP(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin) (bool, jwt.MapClaims,string, error) {
	var (
		ldapServer,ldapPort,ldapBindDN,ldapPassword,ldapSearchDN    string
	)
	if reqLogin.LocationDepartment == "HQ" {
		ldapServer   = os.Getenv("LDAP_SERVER_HQ")
		ldapPort     = os.Getenv("LDAP_PORT_HQ")
		ldapBindDN   = os.Getenv("LDAP_BIND_DN_HQ")
		ldapPassword = os.Getenv("LDAP_PASSWORD_HQ")
		ldapSearchDN = os.Getenv("LDAP_SEARCH_DN_HQ")
	}else if reqLogin.LocationDepartment == "BR" {
		ldapServer   = os.Getenv("LDAP_SERVER_BR")
		ldapPort     = os.Getenv("LDAP_PORT_BR")
		ldapBindDN   = os.Getenv("LDAP_BIND_DN_BR")
		ldapPassword = os.Getenv("LDAP_PASSWORD_BR")
		ldapSearchDN = os.Getenv("LDAP_SEARCH_DN_BR")
	}

	tcpdial :=  fmt.Sprintf("%s:%v", ldapServer, ldapPort)
	l, err := ldap.Dial("tcp", tcpdial)
    if err != nil {
		log.Println("INI HANYA LOG TIDAK MENGELUARKAN APAPUN => ldap.Dial(tcp, log.Sprintf(s:v, ldapServer, ldapPort))")
		log.Println(err.Error())
        return false, nil,"", errors.New("invalid username/password")
    }
    defer l.Close()

	err = l.Bind(ldapBindDN, ldapPassword)
	if err != nil {
		log.Println("l.Bind(ldapBindDN, ldapPassword)")
		log.Println(err.Error())
        return false, nil,"", errors.New("invalid username/password")
	}

	searchRequest := ldap.NewSearchRequest(
		ldapSearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(sAMAccountName=%s)", reqLogin.Username),
		[]string{"dn", "displayName", "department", "whenChanged", "sAMAccountName", "userAccountControl", "mail","promotsrole","lastLogonTimestamp","badPwdCount","physicalDeliveryOfficeName"},
		nil,
	)
	
	sr, err := l.Search(searchRequest)
	
	if err != nil {
		log.Println("err != nil")
		log.Println(err.Error())
		return false, nil,"", errors.New("invalid username/password")
	}

	if len(sr.Entries) == 0 {
		log.Println("len(sr.Entries) == 0")
		return false, nil,"", errors.New("user tidak ditemukan")
	}
	

	entry := sr.Entries[0]
	var adCodeMessage = new([]loginmodel.ADCodeMessage)

	log.Printf("userAccountControl => %v\n", entry.GetAttributeValue("userAccountControl"))
	if entry.GetAttributeValue("userAccountControl") == "" {
		log.Printf("entry.GetAttributeValue(userAccountControl) => %v", entry.GetAttributeValue("userAccountControl"))
		log.Println(err)
        return false, nil,"", errors.New("userAccountControl tidak ditemukan")
	}

	ldapLoginServiceImpl.LoginRepository.ADCodeMessage(f,adCodeMessage,"userAccountControl")
	for _, v := range *adCodeMessage {
		if v.Code == entry.GetAttributeValue("userAccountControl") {
			log.Println("strconv.Itoa(v.Code) == entry.GetAttributeValue(userAccountControl)")
			if !v.IsSuccess {
				// return false, nil,"", errors.New(v.Message +" - Status Code: " + v.Code)
				return false, nil,"", errors.New(v.Message)
			}
		}
	}

	log.Printf("mail => %v\n", entry.GetAttributeValue("mail"))
	if entry.GetAttributeValue("mail") == "" {
		log.Println("entry.GetAttributeValue(mail)")
		log.Println(err)
        return false, nil,"", errors.New("mail tidak ditemukan")
	}
	
	log.Printf("physicalDeliveryOfficeName => %v\n",entry.GetAttributeValue("physicalDeliveryOfficeName"))
	if entry.GetAttributeValue("physicalDeliveryOfficeName") == "" {
		log.Println("entry.GetAttributeValue(physicalDeliveryOfficeName)")
		log.Println(err)
        return false, nil,"", errors.New("Kode Unit null/kosong")
	}

	log.Printf("promotsrole => %v\n",entry.GetAttributeValue("promotsrole"))
	if entry.GetAttributeValue("promotsrole") == "" {
		log.Println("entry.GetAttributeValue(promotsrole)")
		log.Println(err)
        return false, nil,"", errors.New("User tidak memiliki hak akses, Silahkan hubungi Admin")
	}

	ldapLoginServiceImpl.LoginRepository.ADCodeMessage(f,adCodeMessage,"promotsrole")
	var isSuccessPromotsRole = true

	for _, v := range *adCodeMessage {
		if v.Code == entry.GetAttributeValue("promotsrole") {
			log.Println("Code:", v.Code, "matches : ", entry.GetAttributeValue("promotsrole"))
			isSuccessPromotsRole = false
		}
	}

	if isSuccessPromotsRole {
		// return false, nil, "", errors.New(entry.GetAttributeValue("promotsrole") + " tidak ditemukan di sistem kami")
		return false, nil, "", errors.New("Kewenangan tidak terdaftar")
	}
	
	// Unit Role Validation
	unitRole := new([]loginmodel.UnitRole)
	ldapLoginServiceImpl.LoginRepository.UnitRole(f,unitRole,entry.GetAttributeValue("physicalDeliveryOfficeName"),entry.GetAttributeValue("promotsrole"))
	log.Println("promotsrole => "+ entry.GetAttributeValue("promotsrole"))
	log.Println("physicalDeliveryOfficeName - "+entry.GetAttributeValue("physicalDeliveryOfficeName"))
	if len(*unitRole) < 1 {
		// return false, nil, "", errors.New("kode unit: " + entry.GetAttributeValue("physicalDeliveryOfficeName") + " dan promotRole "+ entry.GetAttributeValue("promotsrole") +" tidak ditemukan di sistem kami")
		return false, nil, "", errors.New("Kode unit tidak terdaftar")
	}

	log.Printf("lastLogonTimestamp => %+v\n",entry.GetAttributeValue("lastLogonTimestamp"))
	log.Println(entry.GetAttributeValue("lastLogonTimestamp"))
	if entry.GetAttributeValue("lastLogonTimestamp") == "" {
		log.Println("entry.GetAttributeValue(lastLogonTimestamp)")
		log.Println(err)
        return false, nil,"", errors.New("lastLogonTimestamp tidak ditemukan")
	}

	log.Printf("badPwdCount => %+v\n",entry.GetAttributeValue("badPwdCount"))
	log.Println(entry.GetAttributeValue("badPwdCount"))
	if entry.GetAttributeValue("badPwdCount") == "" {
		log.Println("entry.GetAttributeValue(badPwdCount)")
		log.Println(err)
        return false, nil,"", errors.New("badPwdCount tidak ditemukan")
	}

	badPwdCountInt,_ := strconv.Atoi(entry.GetAttributeValue("badPwdCount"))
	if badPwdCountInt >= 10 {
		log.Println("entry.GetAttributeValue(badPwdCount) 10")
		log.Println(err)
        return false, nil,"", errors.New("User anda Terkunci, Silahkan hubungi Admin")
	}


	userDn := sr.Entries[0].DN
	err = l.Bind(userDn, reqLogin.Password)
	if err != nil {
		log.Println(err.Error())
		return false, nil,"", errors.New("invalid username/password")
	}

	var roleMenuView = new([]loginmodel.RoleMenuView)
	if err := ldapLoginServiceImpl.LoginRepository.RoleMenuView(f,roleMenuView,entry.GetAttributeValue("promotsrole")); err != nil{
		return false, nil,"", errors.New("error role menu view")
	}

	userInfo := map[string]string{
		"displayName":        entry.GetAttributeValue("displayName"),
		"department":         entry.GetAttributeValue("department"),
		// "maverickApps":       entry.GetAttributeValue("maverickApps"),
		"whenChanged":        entry.GetAttributeValue("whenChanged"),
		"sAMAccountName":     entry.GetAttributeValue("sAMAccountName"),
		"userAccountControl": entry.GetAttributeValue("userAccountControl"),
		"userMail":           entry.GetAttributeValue("mail"),
		"promotsRole":        entry.GetAttributeValue("promotsrole"),
		"lastLogonTimestamp":           entry.GetAttributeValue("lastLogonTimestamp"),
		"badPwdCount":           entry.GetAttributeValue("badPwdCount"),
		"physicalDeliveryOfficeName":        entry.GetAttributeValue("physicalDeliveryOfficeName"),
	}
	token,claims, _ := ldapLoginServiceImpl.JWTTokenClaims(f,userInfo)
	
	return true, claims,token, nil
}

func (ldapLoginServiceImpl *LdapLoginServiceImpl) AuthVendor(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin) (string,jwt.MapClaims,error) {
	users := new([]usermodel.Users)
	err := ldapLoginServiceImpl.LoginRepository.CheckUser(f,reqLogin, users)
	var dataUserLogin loginmodel.DataUserLogin
	if err != nil {
		log.Println("ldapLoginServiceImpl.LoginRepository.CheckUser")
		return "",jwt.MapClaims{},err
	}

	for _, user := range *users {
		dataUserLogin.Username = user.CODE
		dataUserLogin.Password = user.PASSWORD
		dataUserLogin.RoleName = append(dataUserLogin.RoleName, user.ROLE_NAME)
	} 
	plainTextPassword,err := decrypt.DecryptAES(dataUserLogin.Password)
	if err != nil {
		return "",jwt.MapClaims{}, errors.New("error DecryptAES")
	}
	
	if reqLogin.Password != plainTextPassword {
		log.Println("password was wrong")
		return "",jwt.MapClaims{}, errors.New("username or password wrong")
	}
	var responseLogin = loginmodel.ResponseDataUserLogin{
		Username: dataUserLogin.Username,
		RoleName: dataUserLogin.RoleName,
	}
	token, dataUserResult, err := ldapLoginServiceImpl.JWTTokenClaims(f,responseLogin)
	if err != nil {
		log.Println("password was wrong")
		return "",jwt.MapClaims{}, nil
	}
	return token, dataUserResult, nil
}	

func (ldapLoginServiceImpl *LdapLoginServiceImpl) JWTTokenClaims(f *fiber.Ctx,data any) (string, jwt.MapClaims, error) {
	time_env := os.Getenv("TIME_JWT_EXP")
	secret_token := os.Getenv("SECRET_TOKEN")
	timeInt,err := strconv.Atoi(time_env)
	if err != nil {
		log.Println("strconv.Atoi(time_env) err")
		return "",jwt.MapClaims{},err
	}
	claims := jwt.MapClaims{
		"user":  data,
		"exp":   time.Now().Add(time.Minute * time.Duration(timeInt)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secret_token))
	if err != nil {
		log.Println("t, err, err")
		return "",jwt.MapClaims{},err
	}

	return t,claims,nil
}

func getDataRequestLogin(c *fiber.Ctx) error {
	// Get IP address
	ipAddress := c.IP()

	// Get MAC address
	macAddress := getMacAddress()

	// Get hostname
	hostname, _ := os.Hostname()

	// Get browser agent
	browserAgent := c.Get("User-Agent")
	// curl -H "User-Agent: MyCustomUserAgent/1.0" http://localhost:3000/testing/login

	// Get OS version
	osVersion, _ := getOSVersion()

	response := fmt.Sprintf("IP Address: %s\nMAC Address: %s\nHostname: %s\nBrowser Agent: %s\nOS Version: %s\n",
		ipAddress, macAddress, hostname, browserAgent, osVersion)
	fmt.Println(response)
	return nil
}

func getMacAddress() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, interf := range interfaces {
		if interf.Flags&net.FlagUp != 0 && !strings.HasPrefix(interf.Name, "lo") {
			return interf.HardwareAddr.String()
		}
	}
	return ""
}

func getOSVersion() (string, error) {
	info, err := host.Info()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s", info.Platform, info.PlatformVersion), nil
}