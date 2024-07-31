package loginrepositories

import (
	"github.com/gofiber/fiber/v2"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
	usermodel "github.com/lenna-ai/bni-iproc/models/userModel"
	"gorm.io/gorm"
)

type LoginRepositoryImpl struct {
	DB *gorm.DB
}

type LoginRepository interface {
	CheckUser(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin,users *[]usermodel.Users)(error)
	ADCodeMessage(f *fiber.Ctx, dataCode *[]loginmodel.ADCodeMessage,attribute string) error
	UnitRole(f *fiber.Ctx, UnitRole *[]loginmodel.UnitRole, physicalDeliveryOfficeName string) error
	RoleMenuView(f *fiber.Ctx, roleMenuView *[]loginmodel.RoleMenuView,roleName string) error
}

func NewLoginRepository(db *gorm.DB) *LoginRepositoryImpl {
	return &LoginRepositoryImpl{
		DB: db,
	}
}