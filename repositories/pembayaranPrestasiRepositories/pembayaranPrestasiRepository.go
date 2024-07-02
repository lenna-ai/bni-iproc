package pembayaranprestasirepositories

import (
	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
	"gorm.io/gorm"
)

type PembayaranPrestasiRepositoryImpl struct {
	DB *gorm.DB
}
type PembayaranPrestasiRepository interface {
	DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi) error
}

func NewPembayaranPrestasiRepository(db *gorm.DB) *PembayaranPrestasiRepositoryImpl {
	return &PembayaranPrestasiRepositoryImpl{
		DB: db,
	}
}
