package dashboardrekananservices

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
)

func (dashboardRekananServiceImpl *DashboardRekananServiceImpl) Rekanan(c *fiber.Ctx,param string,rekananData *[]map[string]any,totalCount *int64) error {
	if err := dashboardRekananServiceImpl.DashboardRekananRepository.Rekanan(c,param,rekananData,totalCount); err != nil{
		return err
	}
	return nil
}
func (dashboardRekananServiceImpl *DashboardRekananServiceImpl) BreakdownRekanan(c *fiber.Ctx,param string,breakdownRekananData *[]dashboardmodel.DashboardRekanan,totalCount *int64) error {
	if err := dashboardRekananServiceImpl.DashboardRekananRepository.BreakdownRekanan(c,param,breakdownRekananData,totalCount); err != nil{
		return err
	}
	return nil
}