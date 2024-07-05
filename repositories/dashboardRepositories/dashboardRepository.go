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
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepositoryImpl {
	return &DashboardRepositoryImpl{
		DB: db,
	}
}