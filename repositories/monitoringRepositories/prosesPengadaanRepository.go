package monitoringrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
	formatterProsesPengadaanModel "github.com/lenna-ai/bni-iproc/models/prosesPengadaanModel/formatters"
	"gorm.io/gorm"
)

type MonitoringProsesPengadaanImpl struct {
	DB *gorm.DB
}

type MonitoringProsesPengadaan interface {
	JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error)
	GetProsesPengadaan(c *fiber.Ctx,totalCount *int64) (*[]formatterProsesPengadaanModel.PutPengadaanFormatter, error)
	PutProsesPengadaan(c *fiber.Ctx, prosesPengadaanModel *formatterProsesPengadaanModel.PutPengadaanFormatter) error
}

func NewMonitoringProsesPengadaan(db *gorm.DB) *MonitoringProsesPengadaanImpl {
	return &MonitoringProsesPengadaanImpl{
		DB: db,
	}
}
