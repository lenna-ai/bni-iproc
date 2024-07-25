package dashboardrekananrepositories

import (
	"github.com/gofiber/fiber/v2"
	pegadaanmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
	"gorm.io/gorm"
)

type DashboardRekananRepository interface {
	Rekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaVendor string,rekananData *[]map[string]any,totalCount *int64) error
	BreakdownRekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaPekerjaan string,breakdownRekananData *[]pegadaanmodel.PengadaanFilter, totalCount *int64) error
}

type DashboardRekananRepositoryImpl struct {
	DB *gorm.DB
}

func NewDashboardRekananRepository(db *gorm.DB) *DashboardRekananRepositoryImpl {
	return &DashboardRekananRepositoryImpl{
		DB: db,
	}
}