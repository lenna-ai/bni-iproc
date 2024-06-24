package monitoringrepositories

import (
	"github.com/gofiber/fiber/v2"
	pegadaanmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
)

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error) {
	var jenisPengadaan = new([]formatters.JenisPengadaan)
	var pengadaanModel = new([]pegadaanmodel.Pengadaan)
	if err := monitoringProsesPengadaanImpl.DB.Model(pengadaanModel).Select("JENIS_PENGADAAN").Group("JENIS_PENGADAAN").Find(&jenisPengadaan).Error; err != nil {
		return jenisPengadaan, err
	}
	return jenisPengadaan, nil
}
