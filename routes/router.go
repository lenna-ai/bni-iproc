package routes

import (
	"github.com/gofiber/fiber/v2"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
)

func Router(app *fiber.App) {
	PengadaanDoneController := appconfig.InitControllerServiceRepository()
	pengadaan := app.Group("pengadaan")
	pengadaan.Get("/", PengadaanDoneController.IndexPengadaan)
	pengadaan.Get("/filter", PengadaanDoneController.FilterPengadaan)
	pengadaan.Get("/sum", PengadaanDoneController.SumPengadaan)

	pengadaan.Get("/status", PengadaanDoneController.IndexStatus)
	pengadaan.Get("/type_pengadaan", PengadaanDoneController.IndexType)

}
