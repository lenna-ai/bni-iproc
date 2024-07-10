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