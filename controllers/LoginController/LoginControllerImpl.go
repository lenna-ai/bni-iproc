package logincontroller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
)

func (loginControllerImpl *LoginControllerImpl) Ldap(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	reqLogin := new(loginmodel.RequestLogin)
	if err := c.BodyParser(reqLogin); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	isSuccess, data,token, err := loginControllerImpl.LdapLoginService.AuthUsingLDAP(c, reqLogin);
	if !isSuccess {
		return helpers.ResultUnauthorizedJsonApi(c, nil, errors.New("invalid username/password").Error())
	}
	if  err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	result := fiber.Map{
		"data":data,
		"token":token,
	}
	
	return helpers.ResultSuccessJsonApi(c,result)
}