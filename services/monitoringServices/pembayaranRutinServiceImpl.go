package monitoringservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
)

func (pembayaranRutinServiceImpl *PembayaranRutinServiceImpl) DetailPembayaranRutin(c *fiber.Ctx, pembayaranRutin *[]pembayaranrutinmodel.PembayaranRutin) error {
	return pembayaranRutinServiceImpl.PembayaranRutinRepository.DetailPembayaranRutin(c, pembayaranRutin)
}
func (pembayaranRutinServiceImpl *PembayaranRutinServiceImpl) PutPembayaranRutin(c *fiber.Ctx, pembayaranRutin *pembayaranrutinmodel.PembayaranRutin) error {
	return pembayaranRutinServiceImpl.PembayaranRutinRepository.PutPembayaranRutin(c, pembayaranRutin)
}
