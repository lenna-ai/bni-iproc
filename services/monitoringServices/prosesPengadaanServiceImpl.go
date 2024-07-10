package monitoringservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
	formatterProsesPengadaanModel "github.com/lenna-ai/bni-iproc/models/prosesPengadaanModel/formatters"
)

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error) {
	jenisPengadaan, err := monitoringProsesPengadaanImpl.monitoringRepository.JenisPengadaan(c)
	if err != nil {
		return jenisPengadaan, nil
	}
	return jenisPengadaan, nil
}
func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) DetailProsesPengadaan(c *fiber.Ctx,totalCount *int64) (*[]formatterProsesPengadaanModel.PutPengadaanFormatter, error) {
	prosesPengadaanModel, err := monitoringProsesPengadaanImpl.monitoringRepository.GetProsesPengadaan(c,totalCount)
	if err != nil {
		return prosesPengadaanModel, err
	}
	return prosesPengadaanModel, nil
}
func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) PutProsesPengadaan(c *fiber.Ctx, prosesPengadaanModel *formatterProsesPengadaanModel.PutPengadaanFormatter) error {
	if err := monitoringProsesPengadaanImpl.monitoringRepository.PutProsesPengadaan(c, prosesPengadaanModel); err != nil {
		return err
	}
	return nil
}
