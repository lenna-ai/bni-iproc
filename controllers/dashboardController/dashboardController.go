package dashboardcontroller

import (
	"github.com/gofiber/fiber/v2"
	dashboardservices "github.com/lenna-ai/bni-iproc/services/dashboardServices"
)

type DashboardControllerImpl struct {
	DashboardService dashboardservices.DashboardService
}

type DashboardController interface {
	TotalPengadaan(c *fiber.Ctx) error
	TotalPembayaran(c *fiber.Ctx) error
	TotalVendor(c *fiber.Ctx) error
	PengadaanOnGoingKewenangan(c *fiber.Ctx) error
}

func NewDashboardController(dashboardservices dashboardservices.DashboardService,) *DashboardControllerImpl {
	return &DashboardControllerImpl{
		DashboardService: dashboardservices,
	}
}