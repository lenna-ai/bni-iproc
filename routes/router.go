package routes

import (
	"github.com/gofiber/fiber/v2"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
)

func Router(app *fiber.App) {
	pengadaanController, pembayaranMonitoringController := appconfig.InitControllerServiceRepository()
	pengadaan := app.Group("pengadaan")
	pengadaan.Get("/", pengadaanController.IndexPengadaan)
	pengadaan.Get("/filter", pengadaanController.FilterPengadaan)
	pengadaan.Get("/sum", pengadaanController.SumPengadaan)
	pengadaan.Get("/efisiensi", pengadaanController.EfisiensiPengadaan)

	pembayaraan := app.Group("pembayaraan")
	pembayaraan.Get("rekanan", pembayaranMonitoringController.IndexPengadaan)

	pengadaan.Get("/status", pengadaanController.IndexStatus)
	pengadaan.Get("/type_pengadaan", pengadaanController.IndexType)

}
