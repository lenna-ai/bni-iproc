package loginservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
	loginrepositories "github.com/lenna-ai/bni-iproc/repositories/loginRepositories"
)

type LdapLoginServiceImpl struct {
	LoginRepository loginrepositories.LoginRepository
}

type LdapLoginService interface {
	AuthUsingLDAP(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin) (bool, jwt.MapClaims,string, error)
	AuthVendor(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin) (string,jwt.MapClaims,error)
	JWTTokenClaims(f *fiber.Ctx,data any) (string, jwt.MapClaims, error)
}

func NewLdapLoginService(loginRepository loginrepositories.LoginRepository) *LdapLoginServiceImpl {
	return &LdapLoginServiceImpl{
		LoginRepository: loginRepository,
	}
}