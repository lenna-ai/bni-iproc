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
func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) DetailProsesPengadaan(c *fiber.Ctx,totalCount *int64,jenis_pengadaan string) (*[]map[string]interface{}, error) {
	var whereQuery string
	if jenis_pengadaan == "it" {
		whereQuery = "where p.JENIS_PENGADAAN in ('IT')"
	}else if jenis_pengadaan == "premises" {	
		whereQuery = "where p.JENIS_PENGADAAN in ('Premises')"
	}else{
		whereQuery = "where p.JENIS_PENGADAAN NOT in ('IT','Premises')"
	}


	prosesPengadaanModel, err := monitoringProsesPengadaanImpl.monitoringRepository.GetProsesPengadaan(c,totalCount,jenis_pengadaan,whereQuery)
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
