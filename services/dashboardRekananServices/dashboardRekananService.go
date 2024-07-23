package dashboardrekananservices

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
	dashboardrekananrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRekananRepositories"
)

type DashboardRekananService interface {
	Rekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaVendor string,rekananData *[]map[string]any,totalCount *int64) error
	BreakdownRekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaPekerjaan string,breakdownRekananData *[]dashboardmodel.DashboardRekanan,totalCount *int64) error
}

type DashboardRekananServiceImpl struct {
	DashboardRekananRepository dashboardrekananrepositories.DashboardRekananRepository
}

func NewDashboardRekananService(DashboardRekananRepository dashboardrekananrepositories.DashboardRekananRepository) *DashboardRekananServiceImpl {
	return &DashboardRekananServiceImpl{
		DashboardRekananRepository: DashboardRekananRepository,
	}
}