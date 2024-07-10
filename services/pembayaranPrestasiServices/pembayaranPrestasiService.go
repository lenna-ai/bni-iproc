package pembayaranprestasiservices

import (
	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel/breakdown"
	pembayaranprestasirepositories "github.com/lenna-ai/bni-iproc/repositories/pembayaranPrestasiRepositories"
)

type PembayaranPrestasiServiceImpl struct {
	PembayaranPrestasiRepository pembayaranprestasirepositories.PembayaranPrestasiRepository
}

type PembayaranPrestasiService interface {
	DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi, requestPembayaranPrestasi *pembayaranprestasimodel.RequestPembayaranPrestasi,totalCount *int64) error
	PutPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *pembayaranprestasimodel.PembayaranPrestasi) error
	DetailBreakdownPembayaranPrestasi(c *fiber.Ctx, breakdownPembayaraanPrestasi *[]breakdown.BreakdownPembayaranPrestasi, breakdownRequestBreakdownPembayaranPrestasi *breakdown.RequestBreakdownPembayaranPrestasi,totalCount *int64) error
	PutBreakdownPembayaranPrestasi(c *fiber.Ctx, breakdownRequestPutPembayaraanPrestasi *breakdown.RequestPutBreakdownPembayaranPrestasi) error
}

func NewPembayaranPrestasiService(pembayaranPrestasiRepository pembayaranprestasirepositories.PembayaranPrestasiRepository) *PembayaranPrestasiServiceImpl {
	return &PembayaranPrestasiServiceImpl{
		PembayaranPrestasiRepository: pembayaranPrestasiRepository,
	}
}
