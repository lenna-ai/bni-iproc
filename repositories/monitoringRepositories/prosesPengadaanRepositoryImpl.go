package monitoringrepositories

import (
	"log"
	"time"

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

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) GetProsesPengadaan(c *fiber.Ctx, totalCount *int64, jenis_pengadaan string, whereQuery string) (*[]map[string]interface{}, error) {
	data := new([]map[string]interface{})

	// Membuat subquery dasar tanpa LIMIT dan OFFSET
	subQuery := monitoringProsesPengadaanImpl.DB.
		Select(`NVL(MPPN.ID, 0) as mmpnId,
				p.PROCUREMENT_ID,
				P.JENIS_PENGADAAN,
				P.NAMA,
				P.METODE,
				P.TAHAPAN,
				P.SLA_IN_DAYS,
				TO_CHAR(P.SCHEDULE_START_DATE) as START_DATE,
				TO_CHAR(P.SCHEDULE_END_DATE) as END_DATE,
				MPPN.STATUS,
				MPPN.STATUS_PENGADAAN_PROMOTS,
				MPPN.KETERANGAN_JIKA_TERLAMBAT`).
		Table("PENGADAAN p").
		Joins("LEFT JOIN MONITORING_PROSES_PENGADAAN_NEW mppn ON p.PROCUREMENT_ID = MPPN.PROCUREMENT_ID").
		Where(whereQuery).
		Group(`p.PROCUREMENT_ID,
				P.JENIS_PENGADAAN,
				P.NAMA,
				P.METODE,
				P.TAHAPAN,
				P.SLA_IN_DAYS,
				TO_CHAR(P.SCHEDULE_START_DATE),
				TO_CHAR(P.SCHEDULE_END_DATE),
				MPPN.ID,
				MPPN.STATUS,
				MPPN.STATUS_PENGADAAN_PROMOTS,
				MPPN.KETERANGAN_JIKA_TERLAMBAT`)

	// Membungkus subquery dan membuat query utama
	query := monitoringProsesPengadaanImpl.DB.
		Table("(?) as newtable", subQuery).
		Order("mmpnId desc")

	// Menghitung total jumlah data tanpa pagination
	if err := query.Count(totalCount).Error; err != nil {
		log.Println("Error counting total records: ", err)
		return data, err
	}

	// Pagination
	pageSize, offset := gormhelpers.PaginateRaw(c)
	if err := query.Offset(offset).Limit(pageSize).Scan(data).Error; err != nil {
		log.Println("Error in executing paginated query: ", err)
		return data, err
	}

	return data, nil
}

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) PutProsesPengadaan(c *fiber.Ctx, prosesPengadaanModel *formatterProsesPengadaanModel.PutPengadaanFormatter) error {
	t := time.Now()
	layoutFormat := time.DateTime
	dataTempProsesPengadaanModel := new([]formatterProsesPengadaanModel.PutPengadaanFormatter)
	monitoringProsesPengadaanImpl.DB.Find(dataTempProsesPengadaanModel, "PROCUREMENT_ID = ? and DELETED_AT is null", prosesPengadaanModel.PROCUREMENT_ID)
	for _, v := range *dataTempProsesPengadaanModel {
		date, _ := time.Parse(layoutFormat, t.Format(time.DateTime))
		v.DELETED_AT = &date
		v.DELETED_BY = prosesPengadaanModel.DELETED_BY
		monitoringProsesPengadaanImpl.DB.Save(v)
	}
	prosesPengadaanModel.DELETED_BY = nil
	createProsesPengadaanModel := monitoringProsesPengadaanImpl.DB.Create(prosesPengadaanModel)
	if err := createProsesPengadaanModel.Error; err != nil {
		log.Println("createProsesPengadaanModel.Error; err")
		return err
	}
	return nil
}
