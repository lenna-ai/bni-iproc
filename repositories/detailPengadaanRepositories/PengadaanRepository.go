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
	FilterPengadaan(c *fiber.Ctx, filter map[string]string) ([]detailmodel.Pengadaan, error)
	IndexPengadaan(c *fiber.Ctx) ([]detailmodel.Pengadaan, error)
	IndexStatus(c *fiber.Ctx) ([]detailmodel.Status, error)
	IndexType(c *fiber.Ctx) ([]detailmodel.Type, error)
}
