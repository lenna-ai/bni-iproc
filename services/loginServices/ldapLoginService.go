package loginservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
)

type LdapLoginServiceImpl struct {
	
}

type LdapLoginService interface {
	AuthUsingLDAP(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin) (bool, jwt.MapClaims,string, error)
	JWTTokenClaims(f *fiber.Ctx,data any) (string, jwt.MapClaims, error)
}

func NewLdapLoginService() *LdapLoginServiceImpl {
	return &LdapLoginServiceImpl{}
}