package dashboardrepositories

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DashboardRepositoryImpl struct {
	DB *gorm.DB
}

type DashboardRepository interface {
	TotalPengadaan(c *fiber.Ctx,dashboardModel *map[string]interface{}) error
	TotalPembayaran(c *fiber.Ctx,dashboardModel *map[string]interface{}) error
	TotalVendor(c *fiber.Ctx,dashboardModel *map[string]interface{}) error
	PengadaanOnGoingKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error
	PengadaanOnGoingStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error
	PengadaanOnGoingMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	PengadaanOnGoingKeputusan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	PengadaanOnDoneKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error
	PengadaanOnDoneStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error
	PengadaanOnDoneMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	PengadaanOnDoneTrenPengadaan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	InformasiRekanan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	DataInformasiRekanan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{
		DB: db,
	}
}