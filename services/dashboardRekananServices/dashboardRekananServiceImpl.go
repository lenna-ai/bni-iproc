package dashboardrekananservices

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
)

func (dashboardRekananServiceImpl *DashboardRekananServiceImpl) Rekanan(c *fiber.Ctx,param string,rekananData *[]map[string]any) error {
	if err := dashboardRekananServiceImpl.DashboardRekananRepository.Rekanan(c,param,rekananData); err != nil{
		return err
	}
	return nil
}
func (dashboardRekananServiceImpl *DashboardRekananServiceImpl) BreakdownRekanan(c *fiber.Ctx,param string,breakdownRekananData *[]dashboardmodel.DashboardRekanan) error {
	if err := dashboardRekananServiceImpl.DashboardRekananRepository.BreakdownRekanan(c,param,breakdownRekananData); err != nil{
		return err
	}
	return nil
}