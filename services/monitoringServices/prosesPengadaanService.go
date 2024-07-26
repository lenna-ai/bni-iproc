package monitoringservices

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"

	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
	formatterProsesPengadaanModel "github.com/lenna-ai/bni-iproc/models/prosesPengadaanModel/formatters"
	monitoringrepositories "github.com/lenna-ai/bni-iproc/repositories/monitoringRepositories"
)

type MonitoringProsesPengadaanImpl struct {
	monitoringRepository monitoringrepositories.MonitoringProsesPengadaan
}

type MonitoringProsesPengadaan interface {
	JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error)
	DetailProsesPengadaan(c *fiber.Ctx,totalCount *int64,jenis_pengadaan string) (*[]map[string]interface{}, error)
	PutProsesPengadaan(c *fiber.Ctx, prosesPengadaanModel *formatterProsesPengadaanModel.PutPengadaanFormatter) error
}

func NewMonitoringProsesPengadaan(monitoringRepository monitoringrepositories.MonitoringProsesPengadaan) *MonitoringProsesPengadaanImpl {
	return &MonitoringProsesPengadaanImpl{
		monitoringRepository: monitoringRepository,
	}
}
