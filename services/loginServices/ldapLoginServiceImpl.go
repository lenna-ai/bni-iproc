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
		ldapServer   = os.Getenv("LDAP_SERVER")
		ldapPort     = os.Getenv("LDAP_PORT")
		ldapBindDN   = os.Getenv("LDAP_BIND_DN")
		ldapPassword = os.Getenv("LDAP_PASSWORD")
		ldapSearchDN = os.Getenv("LDAP_SEARCH_DN")
	)

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
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", reqLogin.Username),
		[]string{"uid", "cn", "sn", "mail","telephoneNumber"},
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
	entry := sr.Entries[0]

	// verify user password by binding to user dn (with user password)
	err = l.Bind(entry.DN, reqLogin.Password)
	if err != nil {
		log.Println("l.Bind(entry.DN, reqLogin.Password) == nil")
		return false, nil,"", err
	}

	// (optional) store data
	data := new(loginmodel.UserLDAPData)
	data.ID = reqLogin.Username

	for _, attr := range entry.Attributes {
		switch attr.Name {
		case "sn":
			data.Name = attr.Values[0]
		case "mail":
			data.Email = attr.Values[0]
		case "cn":
			data.FullName = attr.Values[0]
		case "telephoneNumber":
			data.TelephoneNumber = attr.Values[0]
		}
	}
	token,claims, _ := ldapLoginServiceImpl.JWTTokenClaims(f,data)
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
	token, dataUserResult, err := ldapLoginServiceImpl.JWTTokenClaims(f,dataUserLogin)
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
