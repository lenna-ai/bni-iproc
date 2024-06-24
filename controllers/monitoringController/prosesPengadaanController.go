package monitoringcontroller

import (
	"github.com/gofiber/fiber/v2"
	monitoringservices "github.com/lenna-ai/bni-iproc/services/monitoringServices"
)

type MonitoringProsesPengadaanImpl struct {
	MonitoringProsesPengadaan monitoringservices.MonitoringProsesPengadaan
}

type MonitoringProsesPengadaan interface {
	JenisPengadaan(c *fiber.Ctx) error
}

func NewMonitoringProsesPengadaan(monitoringProsesPengadaan monitoringservices.MonitoringProsesPengadaan) *MonitoringProsesPengadaanImpl {
	return &MonitoringProsesPengadaanImpl{
		MonitoringProsesPengadaan: monitoringProsesPengadaan,
	}
}
