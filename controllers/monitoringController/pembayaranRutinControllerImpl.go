package monitoringcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
)

func (pembayaranRutinControllerImpl *PembayaranRutinControllerImpl) DetailPembayaranRutin(c *fiber.Ctx) error {
	putPembayaranRutinModelModel := new([]pembayaranrutinmodel.PembayaranRutin)
	err := pembayaranRutinControllerImpl.PembayaranRutinService.DetailPembayaranRutin(c, putPembayaranRutinModelModel)
	if err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, putPembayaranRutinModelModel)
}

func (pembayaranRutinControllerImpl *PembayaranRutinControllerImpl) PutPembayaranRutin(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	putPembayaranRutinModelModel := new(pembayaranrutinmodel.PembayaranRutin)
	if err := c.BodyParser(putPembayaranRutinModelModel); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	jsonTag, valueErrorTag, valueErrorParam, err := helpers.ValidationFields(putPembayaranRutinModelModel)
	if err != nil {
		return helpers.MessageErrorValidation(c, jsonTag, valueErrorTag, valueErrorParam)
	}

	err = pembayaranRutinControllerImpl.PembayaranRutinService.PutPembayaranRutin(c, putPembayaranRutinModelModel)
	if err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	return helpers.ResultSuccessJsonApi(c, putPembayaranRutinModelModel)
}
