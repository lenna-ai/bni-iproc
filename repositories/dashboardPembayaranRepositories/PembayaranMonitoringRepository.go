package dashboardpembayaranrepositories

import (
	"github.com/gofiber/fiber/v2"
	pembayaranmodel "github.com/lenna-ai/bni-iproc/models/pembayaranModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	"gorm.io/gorm"
)

type PembayaranMonitoringRepositoryImpl struct {
	DB *gorm.DB
}

type PembayaranMonitoringRepository interface {
	IndexRekananPembayaranMonitor(c *fiber.Ctx, jenisPengadaan string,totalCount *int64) ([]formatters.IndexPembayaranMonitor, error)
	FilterPengadaan(c *fiber.Ctx, queryStringWhere string,totalCount *int64) ([]pembayaranmodel.Pembayaran, error)
}
