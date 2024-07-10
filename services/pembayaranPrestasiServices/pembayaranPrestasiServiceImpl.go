package pembayaranprestasiservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel/breakdown"
)

func (pembayaranPrestasiServiceImpl *PembayaranPrestasiServiceImpl) DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi, requestPembayaranPrestasi *pembayaranprestasimodel.RequestPembayaranPrestasi,totalCount *int64) error {
	err := pembayaranPrestasiServiceImpl.PembayaranPrestasiRepository.DetailPembayaranPrestasi(c, pembayaranPrestasi, requestPembayaranPrestasi,totalCount)
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
func (pembayaranPrestasiServiceImpl *PembayaranPrestasiServiceImpl) DetailBreakdownPembayaranPrestasi(c *fiber.Ctx, breakdownPembayaraanPrestasi *[]breakdown.BreakdownPembayaranPrestasi, breakdownRequestBreakdownPembayaranPrestasi *breakdown.RequestBreakdownPembayaranPrestasi,totalCount *int64) error {
	if err := pembayaranPrestasiServiceImpl.PembayaranPrestasiRepository.DetailBreakdownPembayaranPrestasi(c, breakdownPembayaraanPrestasi, breakdownRequestBreakdownPembayaranPrestasi,totalCount); err != nil {
		return err
	}

	return nil
}
func (pembayaranPrestasiServiceImpl *PembayaranPrestasiServiceImpl) PutBreakdownPembayaranPrestasi(c *fiber.Ctx, breakdownRequestPutPembayaraanPrestasi *breakdown.RequestPutBreakdownPembayaranPrestasi) error {
	if err := pembayaranPrestasiServiceImpl.PembayaranPrestasiRepository.PutBreakdownPembayaranPrestasi(c, breakdownRequestPutPembayaraanPrestasi); err != nil {
		return err
	}
	return nil
}
