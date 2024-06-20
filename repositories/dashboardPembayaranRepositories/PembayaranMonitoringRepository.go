package dashboardpembayaranrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	"gorm.io/gorm"
)

type PembayaranMonitoringRepositoryImpl struct {
	DB *gorm.DB
}

type PembayaranMonitoringRepository interface {
	IndexRekananPembayaranMonitor(c *fiber.Ctx, jenisPengadaan string) ([]formatters.IndexPembayaranMonitor, error)
}
