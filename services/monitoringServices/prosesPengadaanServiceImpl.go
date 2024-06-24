package monitoringservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
)

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error) {
	jenisPengadaan, err := monitoringProsesPengadaanImpl.monitoringRepository.JenisPengadaan(c)
	if err != nil {
		return jenisPengadaan, nil
	}
	return jenisPengadaan, nil
}
