package dashboardrekanancontroller

import (
	"net/url"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	pegadaanmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
)

func (dashboardRekananImpl *DashboardRekananImpl) Rekanan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	var rekananData = new([]map[string]any)
	var totalCount = new(int64)
	var filterNamaVendor string

	param, err := url.QueryUnescape(c.Params("jenis_pengadaan"))
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	filterNamaVendor = c.Query("filterNamaVendor")

	var pagination bool 
	if page != 0 && pageSize != 0 {
		pagination = true
	}

	if err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	if err :=dashboardRekananImpl.DashboardRekananService.Rekanan(c,pagination,param,filterNamaVendor,rekananData,totalCount); err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	
	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,rekananData))
}

func (dashboardRekananImpl *DashboardRekananImpl) BreakdownRekanan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	var dashboardRekananData = new([]pegadaanmodel.PengadaanFilter)
	var totalCount = new(int64)
	var filterNamaPekerjaan string
	var pagination bool 


	param, err := url.QueryUnescape(c.Params("nama_pt"))
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if page != 0 && pageSize != 0 {
		pagination = true
	}
	filterNamaPekerjaan = c.Query("filterNamaPekerjaan")

	if err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}
	if err :=dashboardRekananImpl.DashboardRekananService.BreakdownRekanan(c,pagination,param,filterNamaPekerjaan,dashboardRekananData, totalCount); err != nil {
		helpers.ResultFailedJsonApi(c,nil,err.Error())
	}

	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,dashboardRekananData))
}