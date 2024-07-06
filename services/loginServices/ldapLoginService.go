package loginservices

import (
	"github.com/gofiber/fiber/v2"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
)

type LdapLoginServiceImpl struct {
	
}

type LdapLoginService interface {
	AuthUsingLDAP(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin) (bool, *loginmodel.UserLDAPData,string, error)
	JWTTokenClaims(f *fiber.Ctx,data any) (string,error)
}

func NewLdapLoginService() *LdapLoginServiceImpl {
	return &LdapLoginServiceImpl{}
}