package dashboardrekananservices

import (
	"github.com/gofiber/fiber/v2"
	pegadaanmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
	dashboardrekananrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRekananRepositories"
)

type DashboardRekananService interface {
	Rekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaVendor string,rekananData *[]map[string]any,totalCount *int64) error
	BreakdownRekanan(c *fiber.Ctx,usePagination bool,param string,jenis_pengadaan string,filterNamaPekerjaan string,breakdownRekananData *[]pegadaanmodel.PengadaanFilter,totalCount *int64) error
}

type DashboardRekananServiceImpl struct {
	DashboardRekananRepository dashboardrekananrepositories.DashboardRekananRepository
}

func NewDashboardRekananService(DashboardRekananRepository dashboardrekananrepositories.DashboardRekananRepository) *DashboardRekananServiceImpl {
	return &DashboardRekananServiceImpl{
		DashboardRekananRepository: DashboardRekananRepository,
	}
}