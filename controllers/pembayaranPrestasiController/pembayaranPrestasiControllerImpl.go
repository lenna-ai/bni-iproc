package pembayaranprestasicontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
)

func (pembayaranPrestasiControllerImpl *PembayaranPrestasiControllerImpl) DetailPembayaranPrestasi(c *fiber.Ctx) error {
	var pembayaraanPrestasi = new([]pembayaranprestasimodel.PembayaranPrestasi)
	err := pembayaranPrestasiControllerImpl.PembayaranPrestasiService.DetailPembayaranPrestasi(c, pembayaraanPrestasi)
	if err != nil {
		return helpers.ResultFailedJsonApi(c, pembayaraanPrestasi, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, pembayaraanPrestasi)
}
