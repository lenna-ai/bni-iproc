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
	Anggaran(c *fiber.Ctx,anggaran string,anggaranModel *[]map[string]interface{}) error
	PengadaanOnGoingKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error
	PengadaanOnGoingStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error
	PengadaanOnGoingMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	PengadaanOnGoingKeputusan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	PengadaanOnDoneKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error
	PengadaanOnDoneStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error
	PengadaanOnDoneMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	PengadaanOnDoneTrenPengadaan(c *fiber.Ctx,status string,year string,PengadaanOnDoneTrenPengadaan *[]map[string]interface{}) error
	InformasiRekanan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
	DataInformasiRekanan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error
}

func NewDashboardService(dashboardRepository dashboardrepositories.DashboardRepository) *DashboardServiceImpl {
	return &DashboardServiceImpl{
		DashboardRepository: dashboardRepository,
	}
}