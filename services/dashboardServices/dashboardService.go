package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
	dashboardrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRepositories"
)

type DashboardServiceImpl struct {
	DashboardRepository dashboardrepositories.DashboardRepository
}

type DashboardService interface {
	TotalPengadaan(c *fiber.Ctx,dashboardModel *map[string]interface{}) error
	TotalPembayaran(c *fiber.Ctx,dashboardModel *map[string]interface{}) error
	TotalVendor(c *fiber.Ctx,dashboardModel *map[string]interface{}) error
	PengadaanOnGoingKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error
	Status(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error
	Metode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
}

func NewDashboardService(dashboardRepository dashboardrepositories.DashboardRepository) *DashboardServiceImpl {
	return &DashboardServiceImpl{
		DashboardRepository: dashboardRepository,
	}
}