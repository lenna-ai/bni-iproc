package dashboardrekananrepositories

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
	"gorm.io/gorm"
)

type DashboardRekananRepository interface {
	Rekanan(c *fiber.Ctx,param string,rekananData *[]map[string]any,totalCount *int64) error
	BreakdownRekanan(c *fiber.Ctx,param string,breakdownRekananData *[]dashboardmodel.DashboardRekanan, totalCount *int64) error
}

type DashboardRekananRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRekananRepository(db *gorm.DB) *DashboardRekananRepositoryImpl {
	return &DashboardRekananRepositoryImpl{
		DB: db,
	}
}