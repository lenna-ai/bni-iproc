package monitoringcontroller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
)

func (pembayaranRutinControllerImpl *PembayaranRutinControllerImpl) DetailPembayaranRutin(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)

	var totalCount = new(int64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	putPembayaranRutinModelModel := new([]pembayaranrutinmodel.PembayaranRutin)
	err := pembayaranRutinControllerImpl.PembayaranRutinService.DetailPembayaranRutin(c, putPembayaranRutinModelModel,totalCount)
	if err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,putPembayaranRutinModelModel))
	// return helpers.ResultSuccessJsonApi(c, putPembayaranRutinModelModel)
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

func (pembayaranRutinControllerImpl *PembayaranRutinControllerImpl) DetailBreakdownPembayaranRutin(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)

	var totalCount = new(int64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	DetailBreakdownPembayaranRutin := new([]pembayaranrutinmodel.BreakdownPembayaranRutin)
	err := pembayaranRutinControllerImpl.PembayaranRutinService.DetailBreakdownPembayaranRutin(c, DetailBreakdownPembayaranRutin, totalCount)
	if err != nil {
		log.Println("err.Error()")
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,DetailBreakdownPembayaranRutin))
	// return helpers.ResultSuccessJsonApi(c, DetailBreakdownPembayaranRutin)
}

func (pembayaranRutinControllerImpl *PembayaranRutinControllerImpl) PutBreakdownPembayaranRutin(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	PutBreakdownPembayaranRutin := new(pembayaranrutinmodel.BreakdownPembayaranRutin)
	if err := c.BodyParser(PutBreakdownPembayaranRutin); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	jsonTag, valueErrorTag, valueErrorParam, err := helpers.ValidationFields(PutBreakdownPembayaranRutin)
	if err != nil {
		return helpers.MessageErrorValidation(c, jsonTag, valueErrorTag, valueErrorParam)
	}

	err = pembayaranRutinControllerImpl.PembayaranRutinService.PutBreakdownPembayaranRutin(c, PutBreakdownPembayaranRutin)
	if err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	return helpers.ResultSuccessJsonApi(c, PutBreakdownPembayaranRutin)
}
