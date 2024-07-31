package loginrepositories

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	loginmodel "github.com/lenna-ai/bni-iproc/models/loginModel"
	usermodel "github.com/lenna-ai/bni-iproc/models/userModel"
)

func (loginRepositoryImpl LoginRepositoryImpl) CheckUser(f *fiber.Ctx,reqLogin *loginmodel.RequestLogin,users *[]usermodel.Users) (error) {
	findUser := loginRepositoryImpl.DB.Find(users,"CODE = ?", reqLogin.Username)
	if findUser.RowsAffected <= 0 {
		log.Println("Data Not Found")
		return errors.New("username or password wrong")
	}
	if err := findUser.Error; err != nil {
		log.Println("errors")
		log.Println(err.Error())
		return err
	}
	return nil
}

func (loginRepositoryImpl LoginRepositoryImpl) ADCodeMessage(f *fiber.Ctx, dataCode *[]loginmodel.ADCodeMessage,attribute string) error {
	if err := loginRepositoryImpl.DB.Where("ATTRIBUTE = ?",attribute).Find(dataCode).Error; err != nil {
		log.Println("loginRepositoryImpl.DB.Find(dataCode).Error")
		return errors.New(err.Error())
	}
	return nil
}

func (loginRepositoryImpl LoginRepositoryImpl) UnitRole(f *fiber.Ctx, unitRole *[]loginmodel.UnitRole,physicalDeliveryOfficeName string, roleName string) error {
	if err := loginRepositoryImpl.DB.Where("KODE_UNIT = ? and GROUP_ROLE_NAME = ?",physicalDeliveryOfficeName, roleName).Find(unitRole).Error; err != nil {
		log.Println("loginRepositoryImpl.DB.Find(dataCode).Error")
		return errors.New(err.Error())
	}
	return nil
}

func (loginRepositoryImpl LoginRepositoryImpl) RoleMenuView(f *fiber.Ctx, roleMenuView *[]loginmodel.RoleMenuView,roleName string) error {
	if err := loginRepositoryImpl.DB.Where("ROLE_NAME = ?",roleName).Find(roleMenuView).Error; err != nil {
		log.Println("loginRepositoryImpl.DB.Find(dataCode).Error")
		return errors.New(err.Error())
	}
	return nil
}
