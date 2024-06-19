package dashboardpembayaranrepositories

import (
	"log"

	"github.com/gofiber/fiber/v2"
	pembayaranmodel "github.com/lenna-ai/bni-iproc/models/pembayaranModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	"gorm.io/gorm"
)

func NewDashboardMonitoringRepository(db *gorm.DB) *PembayaranMonitoringRepositoryImpl {
	return &PembayaranMonitoringRepositoryImpl{
		DB: db,
	}
}

func (pembayaranMonitoringRepositoryImpl *PembayaranMonitoringRepositoryImpl) IndexPembayaranMonitor(c *fiber.Ctx, jenisPengadaan string) ([]formatters.IndexPembayaranMonitor, error) {
	var pembayaran = new([]formatters.IndexPembayaranMonitor)
	if err := pembayaranMonitoringRepositoryImpl.DB.Model(&pembayaranmodel.Pembayaran{}).Select("NAMA_VENDOR, COUNT(NAMA_VENDOR) AS total_pekerjaan, sum(NILAI_KONTRAK) AS total_nilai_kontrak").Where("JENIS_PENGADAAN = ?", jenisPengadaan).Group("NAMA_VENDOR").Find(pembayaran).Error; err != nil {
		log.Printf("error pembayaranMonitoringRepositoryImpl.DB.Model %v\n ", err)
		return *pembayaran, err
	}
	return *pembayaran, nil
}
