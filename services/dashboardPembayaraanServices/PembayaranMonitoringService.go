package dashboardpembayaranservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
)

type PembayaranMonitoringServicesImpl struct {
	PembayaranMonitoringRepository dashboardpembayaranrepositories.PembayaranMonitoringRepository
}

type PembayaranMonitoringServices interface {
	IndexPengadaanService(c *fiber.Ctx, jenisPengadaan string) ([]formatters.IndexPembayaranMonitor, error)
}
