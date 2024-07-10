package monitoringrepositories

import (
	"github.com/gofiber/fiber/v2"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
	"gorm.io/gorm"
)

type PembayaranRutinRepository interface {
	DetailPembayaranRutin(c *fiber.Ctx, pembayaranRutin *[]pembayaranrutinmodel.PembayaranRutin,totalCount *int64) error
	PutPembayaranRutin(c *fiber.Ctx, pembayaranRutin *pembayaranrutinmodel.PembayaranRutin) error
	DetailBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *[]pembayaranrutinmodel.BreakdownPembayaranRutin,totalCount *int64) error
	PutBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *pembayaranrutinmodel.BreakdownPembayaranRutin) error
}

type PembayaranRutinRepositoryImpl struct {
	DB *gorm.DB
}

func NewPembayaranRutinRepository(db *gorm.DB) *PembayaranRutinRepositoryImpl {
	return &PembayaranRutinRepositoryImpl{
		DB: db,
	}
}
