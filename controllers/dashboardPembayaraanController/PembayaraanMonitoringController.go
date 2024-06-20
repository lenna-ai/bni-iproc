package dashboardpembayarancontroller

import (
	"github.com/gofiber/fiber/v2"
	dashboardpembayaraanservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
)

type PembayaranMonitoringControllerImpl struct {
	PembayaranMonitoringServices dashboardpembayaraanservices.PembayaranMonitoringServices
}

type PembayaranMonitoringController interface {
	IndexRekananPembayaran(c *fiber.Ctx) error
	FilterPengadaan(c *fiber.Ctx) error
}
