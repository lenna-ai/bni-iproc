package loginController

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	jwthelpers "github.com/lenna-ai/bni-iproc/helpers/jwtHelpers"
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
		return helpers.ResultUnauthorizedJsonApi(c, nil, errors.New("invalid username/password").Error())
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

	token,data,err := loginControllerImpl.LdapLoginService.AuthVendor(c,reqLogin)
	if err != nil {
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