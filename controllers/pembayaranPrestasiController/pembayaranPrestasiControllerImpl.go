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

func (pembayaranPrestasiControllerImpl *PembayaranPrestasiControllerImpl) PutPembayaranPrestasi(c *fiber.Ctx) error {
	var pembayaraanPrestasi = new(pembayaranprestasimodel.PembayaranPrestasi)
	if err := c.BodyParser(pembayaraanPrestasi); err != nil {
		return helpers.ResultFailedJsonApi(c, pembayaraanPrestasi, err.Error())
	}
	jsonTag, valueErrorTag, valueErrorParam, err := helpers.ValidationFields(pembayaraanPrestasi)
	if err != nil {
		return helpers.MessageErrorValidation(c, jsonTag, valueErrorTag, valueErrorParam)
	}
	if err := pembayaranPrestasiControllerImpl.PembayaranPrestasiService.PutPembayaranPrestasi(c, pembayaraanPrestasi); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, pembayaraanPrestasi)
}
