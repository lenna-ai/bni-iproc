package pembayaranprestasiservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
)

func (pembayaranPrestasiServiceImpl *PembayaranPrestasiServiceImpl) DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi) error {
	err := pembayaranPrestasiServiceImpl.PembayaranPrestasiRepository.DetailPembayaranPrestasi(c, pembayaranPrestasi)
	if err != nil {
		return err
	}
	return nil
}
