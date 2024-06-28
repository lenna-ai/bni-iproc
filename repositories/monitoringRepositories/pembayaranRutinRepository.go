package monitoringrepositories

import (
	"github.com/gofiber/fiber/v2"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
	"gorm.io/gorm"
)

type PembayaranRutinRepository interface {
	PutPembayaranRutin(c *fiber.Ctx, pembayaranRutin *pembayaranrutinmodel.PembayaranRutin) error
}

type PembayaranRutinRepositoryImpl struct {
	DB *gorm.DB
}

func NewPembayaranRutinRepository(db *gorm.DB) *PembayaranRutinRepositoryImpl {
	return &PembayaranRutinRepositoryImpl{
		DB: db,
	}
}
