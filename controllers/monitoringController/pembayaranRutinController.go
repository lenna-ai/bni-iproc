package monitoringcontroller

import (
	"github.com/gofiber/fiber/v2"
	monitoringservices "github.com/lenna-ai/bni-iproc/services/monitoringServices"
)

type PembayaranRutinControllerImpl struct {
	PembayaranRutinService monitoringservices.PembayaranRutinService
}

type PembayaranRutinController interface {
	DetailPembayaranRutin(c *fiber.Ctx) error
	PutPembayaranRutin(c *fiber.Ctx) error
	DetailBreakdownPembayaranRutin(c *fiber.Ctx) error
	PutBreakdownPembayaranRutin(c *fiber.Ctx) error
}

func NewPembayaranRutinController(pembayaranRutinService monitoringservices.PembayaranRutinService) *PembayaranRutinControllerImpl {
	return &PembayaranRutinControllerImpl{
		PembayaranRutinService: pembayaranRutinService,
	}
}
