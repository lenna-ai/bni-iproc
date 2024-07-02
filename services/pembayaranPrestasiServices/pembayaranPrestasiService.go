package pembayaranprestasiservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
	pembayaranprestasirepositories "github.com/lenna-ai/bni-iproc/repositories/pembayaranPrestasiRepositories"
)

type PembayaranPrestasiServiceImpl struct {
	PembayaranPrestasiRepository pembayaranprestasirepositories.PembayaranPrestasiRepository
}

type PembayaranPrestasiService interface {
	DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi) error
}

func NewPembayaranPrestasiService(pembayaranPrestasiRepository pembayaranprestasirepositories.PembayaranPrestasiRepository) *PembayaranPrestasiServiceImpl {
	return &PembayaranPrestasiServiceImpl{
		PembayaranPrestasiRepository: pembayaranPrestasiRepository,
	}
}
