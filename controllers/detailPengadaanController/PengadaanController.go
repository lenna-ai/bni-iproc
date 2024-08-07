package detailpengadaancontroller

import (
	"github.com/gofiber/fiber/v2"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
)

type PengadaanControllerImpl struct {
	PengadaanFilterService detailpengadaanservices.PengadaanService
}

type PengadaanDoneController interface {
	IndexPengadaan(c *fiber.Ctx) error
	IndexStatus(c *fiber.Ctx) error
	IndexType(c *fiber.Ctx) error
	FilterPengadaan(c *fiber.Ctx) error
	SumPengadaan(c *fiber.Ctx) error
	EfisiensiPengadaan(c *fiber.Ctx) error
	DynamicPengadaan(c *fiber.Ctx) error
}

func NewDetailPengadaanController(DetailPengadaanFilterService detailpengadaanservices.PengadaanService) *PengadaanControllerImpl {
	return &PengadaanControllerImpl{
		PengadaanFilterService: DetailPengadaanFilterService,
	}
}
