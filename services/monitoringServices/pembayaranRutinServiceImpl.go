package monitoringservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
)

func (pembayaranRutinServiceImpl *PembayaranRutinServiceImpl) DetailPembayaranRutin(c *fiber.Ctx, pembayaranRutin *[]pembayaranrutinmodel.PembayaranRutin, totalCount *int64) error {
	return pembayaranRutinServiceImpl.PembayaranRutinRepository.DetailPembayaranRutin(c, pembayaranRutin,totalCount)
}
func (pembayaranRutinServiceImpl *PembayaranRutinServiceImpl) PutPembayaranRutin(c *fiber.Ctx, pembayaranRutin *pembayaranrutinmodel.PembayaranRutin) error {
	return pembayaranRutinServiceImpl.PembayaranRutinRepository.PutPembayaranRutin(c, pembayaranRutin)
}

func (pembayaranRutinServiceImpl *PembayaranRutinServiceImpl) DetailBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *[]pembayaranrutinmodel.BreakdownPembayaranRutin,totalCount *int64) error {
	return pembayaranRutinServiceImpl.PembayaranRutinRepository.DetailBreakdownPembayaranRutin(c, breakdownPembayaranRutin,totalCount)
}
func (pembayaranRutinServiceImpl *PembayaranRutinServiceImpl) PutBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *pembayaranrutinmodel.BreakdownPembayaranRutin) error {
	return pembayaranRutinServiceImpl.PembayaranRutinRepository.PutBreakdownPembayaranRutin(c, breakdownPembayaranRutin)
}
