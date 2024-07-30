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

func (loginRepositoryImpl LoginRepositoryImpl) ADCodeMessage(f *fiber.Ctx, dataCode *[]loginmodel.ADCodeMessage) error {
	if err := loginRepositoryImpl.DB.Where("ATTRIBUTE = userAccountControl").Find(dataCode).Error; err != nil {
		log.Println("loginRepositoryImpl.DB.Find(dataCode).Error")
		return errors.New(err.Error())
	}
	return nil
}

func (loginRepositoryImpl LoginRepositoryImpl) UnitRole(f *fiber.Ctx, unitRole *[]loginmodel.UnitRole,physicalDeliveryOfficeName string) error {
	if err := loginRepositoryImpl.DB.Where("KODE_UNIT = ?",physicalDeliveryOfficeName).Find(unitRole).Error; err != nil {
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
