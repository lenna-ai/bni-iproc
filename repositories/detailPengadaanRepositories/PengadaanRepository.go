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
	FilterPengadaan(c *fiber.Ctx, stringWhere string) ([]detailmodel.Pengadaan, error)
	SumPengadaan(c *fiber.Ctx, sumSelectStringDetailPengadaan string) ([]detailmodel.DataResultSumPengadaan, error)
}

func NewDetailPengadaanRepository(db *gorm.DB) *PengadaanRepositoryImpl {
	return &PengadaanRepositoryImpl{
		DB: db,
	}
}
