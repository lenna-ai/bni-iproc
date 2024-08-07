package loginController

import (
	"github.com/gofiber/fiber/v2"
	loginservices "github.com/lenna-ai/bni-iproc/services/loginServices"
)

type LoginControllerImpl struct {
	LdapLoginService loginservices.LdapLoginService
}

type LoginController interface {
	Ldap(c *fiber.Ctx) error
	Vendor(c *fiber.Ctx) error
	MeJwt(c *fiber.Ctx) error
	Arifin(c *fiber.Ctx) error
	ErrorHandler(c *fiber.Ctx,err error) error
}

func NewloginController(ldapLoginService loginservices.LdapLoginService) *LoginControllerImpl {
	return &LoginControllerImpl{
		LdapLoginService: ldapLoginService,
	}
}