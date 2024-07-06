package dashboardrekanancontroller

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
)

func (dashboardRekananImpl *DashboardRekananImpl) Rekanan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	var rekananData = new([]map[string]any)
	param, err := url.QueryUnescape(c.Params("jenis_pengadaan"))
	if err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	if err :=dashboardRekananImpl.DashboardRekananService.Rekanan(c,param,rekananData); err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	return helpers.ResultSuccessJsonApi(c,rekananData)
}

func (dashboardRekananImpl *DashboardRekananImpl) BreakdownRekanan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	var dashboardRekananData = new([]dashboardmodel.DashboardRekanan)
	param, err := url.QueryUnescape(c.Params("nama_pt"))
	if err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	if err :=dashboardRekananImpl.DashboardRekananService.BreakdownRekanan(c,param,dashboardRekananData); err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	return helpers.ResultSuccessJsonApi(c,dashboardRekananData)
}