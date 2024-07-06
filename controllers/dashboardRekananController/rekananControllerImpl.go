package dashboardrekanancontroller

import (
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
)

func (dashboardRekananImpl *DashboardRekananImpl) Rekanan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	var rekananData = new([]map[string]any)
	var totalCount = new(int64)

	param, err := url.QueryUnescape(c.Params("jenis_pengadaan"))
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	if err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	if err :=dashboardRekananImpl.DashboardRekananService.Rekanan(c,param,rekananData,totalCount); err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	
	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,rekananData))
}

func (dashboardRekananImpl *DashboardRekananImpl) BreakdownRekanan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	var dashboardRekananData = new([]dashboardmodel.DashboardRekanan)
	var totalCount = new(int64)

	param, err := url.QueryUnescape(c.Params("nama_pt"))
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	if err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	if err :=dashboardRekananImpl.DashboardRekananService.BreakdownRekanan(c,param,dashboardRekananData, totalCount); err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}

	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,dashboardRekananData))
}