package monitoringrepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
	"gorm.io/gorm"
)

type MonitoringProsesPengadaanImpl struct {
	DB *gorm.DB
}

type MonitoringProsesPengadaan interface {
	JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error)
}

func NewMonitoringProsesPengadaan(db *gorm.DB) *MonitoringProsesPengadaanImpl {
	return &MonitoringProsesPengadaanImpl{
		DB: db,
	}
}
