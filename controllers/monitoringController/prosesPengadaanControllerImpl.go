package monitoringcontroller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
)

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) JenisPengadaan(c *fiber.Ctx) error {
	jenisPengadaan, err := monitoringProsesPengadaanImpl.MonitoringProsesPengadaan.JenisPengadaan(c)
	if err != nil {
		log.Printf("error monitoringProsesPengadaanImpl.MonitoringProsesPengadaan.JenisPengadaan %v \n ", err)
		return helpers.ResultFailedJsonApi(c, jenisPengadaan, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, jenisPengadaan)
}
