package pembayaranprestasirepositories

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
)

func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi) error {
	if err := pembayaranPrestasiRepositoryImpl.DB.Find(pembayaranPrestasi).Error; err != nil {
		fmt.Println("pembayaranPrestasiRepositoryImpl.DB.Find(pembayaranPrestasi).Error")
		return err
	}
	return nil
}
