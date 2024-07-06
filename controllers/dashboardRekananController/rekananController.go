package dashboardrekanancontroller

import (
	"github.com/gofiber/fiber/v2"
	dashboardrekananservices "github.com/lenna-ai/bni-iproc/services/dashboardRekananServices"
)

type DashboardRekananImpl struct {
	DashboardRekananService dashboardrekananservices.DashboardRekananService
}

type DashboardRekanan interface {
	Rekanan(c *fiber.Ctx) error
	BreakdownRekanan(c *fiber.Ctx) error
}

func NewDashboardRekananController(dashboardRekananService dashboardrekananservices.DashboardRekananService) *DashboardRekananImpl {
	return &DashboardRekananImpl{
		DashboardRekananService: dashboardRekananService,
	}
}