package dashboardservices

import (
	"github.com/gofiber/fiber/v2"
)

func (dashboardServiceImpl *DashboardServiceImpl) TotalPengadaan(c *fiber.Ctx,dashboardModel *map[string]interface{}) error {
	if err := dashboardServiceImpl.DashboardRepository.TotalPengadaan(c,dashboardModel); err != nil {
		return err
	}
	return nil
}

func (dashboardServiceImpl *DashboardServiceImpl) TotalPembayaran(c *fiber.Ctx,dashboardModel *map[string]interface{}) error {
	if err := dashboardServiceImpl.DashboardRepository.TotalPembayaran(c,dashboardModel); err != nil {
		return err
	}
	return nil
}
func (dashboardServiceImpl *DashboardServiceImpl) TotalVendor(c *fiber.Ctx,dashboardModel *map[string]interface{}) error {
	if err := dashboardServiceImpl.DashboardRepository.TotalVendor(c,dashboardModel); err != nil {
		return err
	}
	return nil
}
func (dashboardServiceImpl *DashboardServiceImpl) PengadaanOnGoingKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error {
	if err := dashboardServiceImpl.DashboardRepository.PengadaanOnGoingKewenangan(c,dashboardModel); err != nil {
		return err
	}
	return nil
}