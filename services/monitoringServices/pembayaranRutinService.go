package monitoringservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
	monitoringrepositories "github.com/lenna-ai/bni-iproc/repositories/monitoringRepositories"
)

type PembayaranRutinService interface {
	DetailPembayaranRutin(c *fiber.Ctx, pembayaranRutin *[]pembayaranrutinmodel.PembayaranRutin) error
	PutPembayaranRutin(c *fiber.Ctx, pembayaranRutin *pembayaranrutinmodel.PembayaranRutin) error
	DetailBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *[]pembayaranrutinmodel.BreakdownPembayaranRutin) error
	PutBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *pembayaranrutinmodel.BreakdownPembayaranRutin) error
}

type PembayaranRutinServiceImpl struct {
	PembayaranRutinRepository monitoringrepositories.PembayaranRutinRepository
}

func NewPembayaranRutinService(pembayaranRutinRepository monitoringrepositories.PembayaranRutinRepository) *PembayaranRutinServiceImpl {
	return &PembayaranRutinServiceImpl{
		PembayaranRutinRepository: pembayaranRutinRepository,
	}
}
