package pembayaranprestasicontroller

import (
	"github.com/gofiber/fiber/v2"
	pembayaranprestasiservices "github.com/lenna-ai/bni-iproc/services/pembayaranPrestasiServices"
)

type PembayaranPrestasiControllerImpl struct {
	PembayaranPrestasiService pembayaranprestasiservices.PembayaranPrestasiService
}

type PembayaranPrestasiController interface {
	DetailPembayaranPrestasi(c *fiber.Ctx) error
	PutPembayaranPrestasi(c *fiber.Ctx) error
}

func NewPembayaranPrestasiController(pembayaranPrestasiService pembayaranprestasiservices.PembayaranPrestasiService) *PembayaranPrestasiControllerImpl {
	return &PembayaranPrestasiControllerImpl{
		PembayaranPrestasiService: pembayaranPrestasiService,
	}
}
