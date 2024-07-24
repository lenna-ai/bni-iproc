package loginservices

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lenna-ai/bni-iproc/helpers/jwtHelpers/decrypt"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
	usermodel "github.com/lenna-ai/bni-iproc/models/userModel"
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

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%v", ldapServer, ldapPort))
    if err != nil {
		log.Println("ldap.Dial(tcp, fmt.Sprintf(s:v, ldapServer, ldapPort))")
        return false, nil,"", err
    }
    defer l.Close()

	err = l.Bind(ldapBindDN, ldapPassword)
	if err != nil {
		log.Println("l.Bind(ldapBindDN, ldapPassword)")
		return false, nil,"", err
	}

	searchRequest := ldap.NewSearchRequest(
		ldapSearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(sAMAccountName=%s)", reqLogin.Username),
		[]string{"dn", "displayName", "department", "maverickApps", "whenChanged", "sAMAccountName", "userAccountControl", "mail"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	
	if err != nil {
		log.Println("err != nil")
		return false, nil,"", err
	}

	if len(sr.Entries) == 0 {
		log.Println("len(sr.Entries) == 0")
		return false, nil,"", errors.New("user not found")
	}

	userDn := sr.Entries[0].DN
	err = l.Bind(userDn, ldapPassword)
	if err != nil {
		return false, nil,"", err
	}

	entry := sr.Entries[0]
	if entry.GetAttributeValue("maverickApps") == "" {
		return false, nil,"", err
	}

	if entry.GetAttributeValue("userAccountControl") == "" {
		return false, nil,"", err
	}

	if entry.GetAttributeValue("mail") == "" {
		return false, nil,"", err
	}

	userInfo := map[string]string{
		"displayName":        entry.GetAttributeValue("displayName"),
		"department":         entry.GetAttributeValue("department"),
		"maverickApps":       entry.GetAttributeValue("maverickApps"),
		"whenChanged":        entry.GetAttributeValue("whenChanged"),
		"sAMAccountName":     entry.GetAttributeValue("sAMAccountName"),
		"userAccountControl": entry.GetAttributeValue("userAccountControl"),
		"userMail":           entry.GetAttributeValue("mail"),
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
	plainTextPassword := decrypt.DecryptAES(dataUserLogin.Password)
	
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
