package monitoringservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
	monitoringrepositories "github.com/lenna-ai/bni-iproc/repositories/monitoringRepositories"
)

type MonitoringProsesPengadaanImpl struct {
	monitoringRepository monitoringrepositories.MonitoringProsesPengadaan
}

type MonitoringProsesPengadaan interface {
	JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error)
}

func NewMonitoringProsesPengadaan(monitoringRepository monitoringrepositories.MonitoringProsesPengadaan) *MonitoringProsesPengadaanImpl {
	return &MonitoringProsesPengadaanImpl{
		monitoringRepository: monitoringRepository,
	}
}
