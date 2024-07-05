package dashboardcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
)

func (dashboardControllerImpl *DashboardControllerImpl) TotalPengadaan(c *fiber.Ctx) error  {
	var totalPengadaan = new(map[string]interface{}) 
	if err := dashboardControllerImpl.DashboardService.TotalPengadaan(c, totalPengadaan); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, totalPengadaan)
}

func (dashboardControllerImpl *DashboardControllerImpl) TotalPembayaran(c *fiber.Ctx) error  {
	var totalPembayaran = new(map[string]interface{}) 
	if err := dashboardControllerImpl.DashboardService.TotalPembayaran(c, totalPembayaran); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, totalPembayaran)
}

func (dashboardControllerImpl *DashboardControllerImpl) TotalVendor(c *fiber.Ctx) error  {
	var totalPembayaran = new(map[string]interface{}) 
	if err := dashboardControllerImpl.DashboardService.TotalVendor(c, totalPembayaran); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, totalPembayaran)
}

func (dashboardControllerImpl *DashboardControllerImpl) PengadaanOnGoingKewenangan(c *fiber.Ctx) error  {
	var totalPembayaran = new([]map[string]interface{}) 
	if err := dashboardControllerImpl.DashboardService.PengadaanOnGoingKewenangan(c, totalPembayaran); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, totalPembayaran)
}

func (dashboardControllerImpl *DashboardControllerImpl) PengadaanOnGoingStatus(c *fiber.Ctx) error  {
	var statusPengadaan = new([]map[string]interface{}) 
	if err := dashboardControllerImpl.DashboardService.PengadaanOnGoingStatus(c,statusPengadaan); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, statusPengadaan)
}

func (dashboardControllerImpl *DashboardControllerImpl) PengadaanOnGoingMetode(c *fiber.Ctx) error  {
	var metodePengadaan = new([]map[string]interface{}) 
	if err := dashboardControllerImpl.DashboardService.PengadaanOnGoingMetode(c,metodePengadaan); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, metodePengadaan)
}


func (dashboardControllerImpl *DashboardControllerImpl) PengadaanOnDoneKewenangan(c *fiber.Ctx) error  {
	var totalPembayaran = new([]map[string]interface{}) 
	if err := dashboardControllerImpl.DashboardService.PengadaanOnDoneKewenangan(c, totalPembayaran); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, totalPembayaran)
}
