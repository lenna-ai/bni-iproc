package loginController

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	jwthelpers "github.com/lenna-ai/bni-iproc/helpers/jwtHelpers"
	"github.com/lenna-ai/bni-iproc/helpers/jwtHelpers/decrypt"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
)

func (loginControllerImpl *LoginControllerImpl) Ldap(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	reqLogin := new(loginmodel.RequestLogin)
	if err := c.BodyParser(reqLogin); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	// password,err := decrypt.DecryptAesFrontend(reqLogin.Password)
	// if  err != nil {
	// 	log.Println("loginControllerImpl.LdapLoginService.AuthUsingLDAP2")
	// 	return helpers.ResultFailedJsonApi(c, nil, err.Error())
	// }
	// reqLogin.Password = password
	
	isSuccess, data,token, err := loginControllerImpl.LdapLoginService.AuthUsingLDAP(c, reqLogin);
	if !isSuccess {
		log.Println("loginControllerImpl.LdapLoginService.AuthUsingLDAP")
		return helpers.ResultUnauthorizedJsonApi(c, nil, err.Error())
	}
	if  err != nil {
		log.Println("loginControllerImpl.LdapLoginService.AuthUsingLDAP2")
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	timeExxp := jwthelpers.ExpJwt(c,data["exp"].(int64))
	result := fiber.Map{
		"data":data["user"],
		"token":token,
		"timeExp":timeExxp,
	}
	
	return helpers.ResultSuccessJsonApi(c,result)
}

func (loginControllerImpl *LoginControllerImpl) Vendor(c *fiber.Ctx) error {
	// defer helpers.RecoverPanicContext(c)
	reqLogin := new(loginmodel.RequestLogin)
	if err := c.BodyParser(reqLogin); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	// password,err := decrypt.DecryptAesFrontend(reqLogin.Password)
	// if  err != nil {
	// 	log.Println("loginControllerImpl.LdapLoginService.AuthUsingLDAP2")
	// 	return helpers.ResultFailedJsonApi(c, nil, err.Error())
	// }
	// reqLogin.Password = password

	token,data,err := loginControllerImpl.LdapLoginService.AuthVendor(c,reqLogin)
	if err != nil {
		fmt.Println(err)
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	timeExxp := jwthelpers.ExpJwt(c,data["exp"].(int64))
	result := fiber.Map{
		"data":data["user"],
		"token":token,
		"timeExp":timeExxp,
	}
	
	return helpers.ResultSuccessJsonApi(c,result)
}

func (loginControllerImpl *LoginControllerImpl) MeJwt(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	data := new(map[string]any)
	jwthelpers.MeJwt(c,data)
	// jwthelpers.ExpJwt(c)

	return helpers.ResultSuccessJsonApi(c,data)
}

func (loginControllerImpl *LoginControllerImpl) ErrorHandler(c *fiber.Ctx,err error) error {
	if err != nil {
		return helpers.ResultUnauthorizedJsonApi(c,nil, errors.New("invalid or expired JWT").Error())
	}
	return nil
}
func (loginControllerImpl *LoginControllerImpl) Arifin(c *fiber.Ctx) error {
	// Load the environment variables (make sure to use a library like godotenv or set them manually)
	// Assuming the environment variables are already loaded
	secret_key_login := os.Getenv("SECRET_KEY_LOGIN")
	key := []byte(secret_key_login)

	// The original data you want to encrypt
	origData := []byte("system")

	// Encrypt the original data
	crypted, err := decrypt.AesEncrypt(origData, key)
	if err != nil {
		panic(err)
	}

	fmt.Println(crypted)

	// Base64 encode the encrypted data
	// encoded := base64.StdEncoding.EncodeToString(crypted)
	fmt.Println("Encrypted and Base64 encoded data:", crypted)

	// You can then decrypt it back to verify
	decrypted, err := decrypt.DecryptAesFrontend(crypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypted data:", decrypted)
	return nil
}