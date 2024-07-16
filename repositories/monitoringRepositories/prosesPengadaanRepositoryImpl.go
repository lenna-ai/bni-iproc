package monitoringrepositories

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	pegadaanmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
	formatterProsesPengadaanModel "github.com/lenna-ai/bni-iproc/models/prosesPengadaanModel/formatters"
)

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) JenisPengadaan(c *fiber.Ctx) (*[]formatters.JenisPengadaan, error) {
	var jenisPengadaan = new([]formatters.JenisPengadaan)
	var pengadaanModel = new([]pegadaanmodel.Pengadaan)
	if err := monitoringProsesPengadaanImpl.DB.Model(pengadaanModel).Select("JENIS_PENGADAAN").Group("JENIS_PENGADAAN").Find(&jenisPengadaan).Error; err != nil {
		log.Println("monitoringProsesPengadaanImpl.DB.Model(pengadaanModel).Select(JENIS_PENGADAAN).Group(JENIS_PENGADAAN).Find(&jenisPengadaan).Error")
		return jenisPengadaan, err
	}
	return jenisPengadaan, nil
}

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) GetProsesPengadaan(c *fiber.Ctx,totalCount *int64) (*[]formatterProsesPengadaanModel.PutPengadaanFormatter, error) {
	var prosesPengadaanModel = new([]formatterProsesPengadaanModel.PutPengadaanFormatter)
	
	monitoringProsesPengadaanImpl.DB.Find(prosesPengadaanModel).Count(totalCount)
	if err := monitoringProsesPengadaanImpl.DB.Scopes(gormhelpers.Paginate(c)).Find(prosesPengadaanModel).Error; err != nil {
		log.Println("monitoringProsesPengadaanImpl.DB.Find(prosesPengadaanModel).Error; err")
		return prosesPengadaanModel, err
	}
	return prosesPengadaanModel, nil
}

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) PutProsesPengadaan(c *fiber.Ctx, prosesPengadaanModel *formatterProsesPengadaanModel.PutPengadaanFormatter) error {
	updateProsesPengadaanModel := monitoringProsesPengadaanImpl.DB.Where("NAMA = ? AND METODE = ? AND TAHAPAN = ? AND STATUS = ? AND SCHEDULE_END_DATE = ? AND SCHEDULE_START_DATE = ?", prosesPengadaanModel.Nama, prosesPengadaanModel.Metode,prosesPengadaanModel.Tahapan,prosesPengadaanModel.Status, prosesPengadaanModel.ScheduleEndDate,prosesPengadaanModel.ScheduleStartDate).Updates(prosesPengadaanModel)
	if updateProsesPengadaanModel.RowsAffected < 1 {
		log.Println("updateProsesPengadaanModel.RowsAffected")
		return errors.New("data not found")
	}
	if err := updateProsesPengadaanModel.Error; err != nil {
		log.Println("updateProsesPengadaanModel.Error; err")
		return err
	}
	return nil
}
