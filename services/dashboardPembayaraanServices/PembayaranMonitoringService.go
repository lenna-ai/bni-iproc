package dashboardpembayaranservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranmodel "github.com/lenna-ai/bni-iproc/models/pembayaranModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
)

type PembayaranMonitoringServicesImpl struct {
	PembayaranMonitoringRepository dashboardpembayaranrepositories.PembayaranMonitoringRepository
}

type PembayaranMonitoringServices interface {
	IndexRekananPembayaranService(c *fiber.Ctx, jenisPengadaan string,totalCount *int64) ([]formatters.IndexPembayaranMonitor, error)
	FilterPengadaan(c *fiber.Ctx, jenisPengadaan string,totalCount *int64) ([]pembayaranmodel.Pembayaran, error)
}
