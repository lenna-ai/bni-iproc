package pembayaranprestasicontroller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel/breakdown"
)

func (pembayaranPrestasiControllerImpl *PembayaranPrestasiControllerImpl) DetailPembayaranPrestasi(c *fiber.Ctx) error {
	var pembayaraanPrestasi = new([]pembayaranprestasimodel.PembayaranPrestasi)
	var requestPembayaranPrestasi = new(pembayaranprestasimodel.RequestPembayaranPrestasi)
	jsonTag, valueErrorTag, valueErrorParam, err := helpers.ValidationFields(requestPembayaranPrestasi)
	if err != nil {
		return helpers.MessageErrorValidation(c, jsonTag, valueErrorTag, valueErrorParam)
	}
	err = pembayaranPrestasiControllerImpl.PembayaranPrestasiService.DetailPembayaranPrestasi(c, pembayaraanPrestasi, requestPembayaranPrestasi)
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

func (pembayaranPrestasiControllerImpl *PembayaranPrestasiControllerImpl) DetailBreakdownPembayaranPrestasi(c *fiber.Ctx) error {
	var breakdownPembayaraanPrestasi = new([]breakdown.BreakdownPembayaranPrestasi)
	var breakdownRequestBreakdownPembayaranPrestasi = new(breakdown.RequestBreakdownPembayaranPrestasi)

	if err := c.BodyParser(breakdownRequestBreakdownPembayaranPrestasi); err != nil {
		return helpers.ResultFailedJsonApi(c, breakdownRequestBreakdownPembayaranPrestasi, err.Error())
	}

	jsonTag, valueErrorTag, valueErrorParam, err := helpers.ValidationFields(breakdownRequestBreakdownPembayaranPrestasi)
	if err != nil {
		return helpers.MessageErrorValidation(c, jsonTag, valueErrorTag, valueErrorParam)
	}

	if err := pembayaranPrestasiControllerImpl.PembayaranPrestasiService.DetailBreakdownPembayaranPrestasi(c, breakdownPembayaraanPrestasi, breakdownRequestBreakdownPembayaranPrestasi); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, breakdownPembayaraanPrestasi)
}

func (pembayaranPrestasiControllerImpl *PembayaranPrestasiControllerImpl) PutBreakdownPembayaranPrestasi(c *fiber.Ctx) error {
	fmt.Println("PutBreakdownPembayaranPrestasi")
	return nil
}
