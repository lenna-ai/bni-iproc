package detailpengadaanrepositories

import (
	"github.com/gofiber/fiber/v2"
	detailmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
	"gorm.io/gorm"
)

type PengadaanRepositoryImpl struct {
	DB *gorm.DB
}

type PengadaanRepository interface {
	IndexPengadaan(c *fiber.Ctx) ([]detailmodel.Pengadaan, error)
	IndexStatus(c *fiber.Ctx) ([]detailmodel.Status, error)
	IndexType(c *fiber.Ctx) ([]detailmodel.Type, error)
	FilterPengadaanUmum(c *fiber.Ctx,usePagination bool, stringWhere string,totalCount *int64) ([]detailmodel.PengadaanFilter, error)
	FilterPengadaanMonitoringPengadaan(c *fiber.Ctx,usePagination bool, stringWhere string,totalCount *int64) ([]detailmodel.PengadaanFilter, error)
	SumPengadaan(c *fiber.Ctx, sumSelectStringDetailPengadaan string) ([]detailmodel.DataResultSumPengadaan, error)
	DynamicPengadaan(c *fiber.Ctx,pagination bool,table string,filter map[string]string,stringWhere string, dataResult *[]map[string]any,totalCount *int64) error
}

func NewDetailPengadaanRepository(db *gorm.DB) *PengadaanRepositoryImpl {
	return &PengadaanRepositoryImpl{
		DB: db,
	}
}
