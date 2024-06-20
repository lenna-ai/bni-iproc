package dashboardpembayarancontroller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	dashboardpembayaraanservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
)

func NewPembayaranMonitoringController(pembayaranMonitoringServices dashboardpembayaraanservices.PembayaranMonitoringServices) *PembayaranMonitoringControllerImpl {
	return &PembayaranMonitoringControllerImpl{
		PembayaranMonitoringServices: pembayaranMonitoringServices,
	}
}

func (pembayaranMonitoringControllerImpl *PembayaranMonitoringControllerImpl) IndexRekananPembayaran(c *fiber.Ctx) error {
	jenisPengadaan := c.Query("JENIS_PENGADAAN")

	pembayaran, err := pembayaranMonitoringControllerImpl.PembayaranMonitoringServices.IndexRekananPembayaranService(c, jenisPengadaan)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexStatus %v\n ", err)
		return helpers.ResultFailedJsonApi(c, pembayaran, err.Error())
	}

	return helpers.ResultSuccessJsonApi(c, pembayaran)
}

func (pembayaranMonitoringControllerImpl *PembayaranMonitoringControllerImpl) FilterPengadaan(c *fiber.Ctx) error {
	jenisPengadaan := c.Query("filter")

	pembayaran, err := pembayaranMonitoringControllerImpl.PembayaranMonitoringServices.FilterPengadaan(c, jenisPengadaan)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexStatus %v\n ", err)
		return helpers.ResultFailedJsonApi(c, pembayaran, err.Error())
	}

	return helpers.ResultSuccessJsonApi(c, pembayaran)
}
