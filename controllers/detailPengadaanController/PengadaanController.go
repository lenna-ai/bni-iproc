package detailpengadaancontroller

import (
	"github.com/gofiber/fiber/v2"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
)

type PengadaanControllerImpl struct {
	PengadaanFilterService detailpengadaanservices.PengadaanService
}

type PengadaanDoneController interface {
	FilterPengadaan(c *fiber.Ctx) error
	IndexPengadaan(c *fiber.Ctx) error
	IndexStatus(c *fiber.Ctx) error
	IndexType(c *fiber.Ctx) error
}
