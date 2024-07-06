package dashboardpembayaranrepositories

import (
	"log"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	pembayaranmodel "github.com/lenna-ai/bni-iproc/models/pembayaranModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	"gorm.io/gorm"
)

func NewDashboardMonitoringRepository(db *gorm.DB) *PembayaranMonitoringRepositoryImpl {
	return &PembayaranMonitoringRepositoryImpl{
		DB: db,
	}
}

func (pembayaranMonitoringRepositoryImpl *PembayaranMonitoringRepositoryImpl) IndexRekananPembayaranMonitor(c *fiber.Ctx, jenisPengadaan string,totalCount *int64) ([]formatters.IndexPembayaranMonitor, error) {
	var pembayaran = new([]formatters.IndexPembayaranMonitor)
	pembayaranMonitoringRepositoryImpl.DB.Model(&pembayaranmodel.Pembayaran{}).Select("NAMA_VENDOR, COUNT(NAMA_VENDOR) AS total_pekerjaan, sum(NILAI_KONTRAK) AS total_nilai_kontrak").Where("JENIS_PENGADAAN = ?", jenisPengadaan).Group("NAMA_VENDOR").Count(totalCount)
	if err := pembayaranMonitoringRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Model(&pembayaranmodel.Pembayaran{}).Select("NAMA_VENDOR, COUNT(NAMA_VENDOR) AS total_pekerjaan, sum(NILAI_KONTRAK) AS total_nilai_kontrak").Where("JENIS_PENGADAAN = ?", jenisPengadaan).Group("NAMA_VENDOR").Find(pembayaran).Error; err != nil {
		log.Printf("error pembayaranMonitoringRepositoryImpl.DB.Model %v\n ", err)
		return *pembayaran, err
	}
	return *pembayaran, nil
}
func (pembayaranMonitoringRepositoryImpl *PembayaranMonitoringRepositoryImpl) FilterPengadaan(c *fiber.Ctx, queryStringWhere string,totalCount *int64) ([]pembayaranmodel.Pembayaran, error) {
	var pembayaranModelEntity = new([]pembayaranmodel.Pembayaran)
	pembayaranMonitoringRepositoryImpl.DB.Model(pembayaranModelEntity).Where(queryStringWhere).Count(totalCount)
	if err := pembayaranMonitoringRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Find(pembayaranModelEntity, queryStringWhere).Error; err != nil {
		log.Printf("error pembayaranMonitoringRepositoryImpl.DB.Model %v\n ", err)
		return *pembayaranModelEntity, err
	}
	return *pembayaranModelEntity, nil
}
