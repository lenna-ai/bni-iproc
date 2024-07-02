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
func (pembayaranPrestasiServiceImpl *PembayaranPrestasiServiceImpl) PutPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *pembayaranprestasimodel.PembayaranPrestasi) error {
	if err := pembayaranPrestasiServiceImpl.PembayaranPrestasiRepository.PutPembayaranPrestasi(c, pembayaranPrestasi); err != nil {
		return err
	}
	return nil
}
