package routes

import (
	"github.com/gofiber/fiber/v2"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
)

func Router(app *fiber.App) {
	pengadaanController, pembayaranMonitoringController, monitoringProsesPengadaanImpl := appconfig.InitControllerServiceRepository()
	pengadaan := app.Group("pengadaan")
	pengadaan.Get("/", pengadaanController.IndexPengadaan)
	pengadaan.Get("/filter", pengadaanController.FilterPengadaan)
	pengadaan.Get("/sum", pengadaanController.SumPengadaan)
	pengadaan.Get("/efisiensi", pengadaanController.EfisiensiPengadaan)

	pembayaraan := app.Group("pembayaraan")
	pembayaraan.Get("/", pembayaranMonitoringController.IndexPembayaran)
	pembayaraan.Get("rekanan", pembayaranMonitoringController.IndexRekananPembayaran)
	pembayaraan.Get("filter", pembayaranMonitoringController.FilterPengadaan)

	monitoring := app.Group("monitoring")
	monitoring.Get("/jenis_pengadaan", monitoringProsesPengadaanImpl.JenisPengadaan)
	monitoring.Get("/proses_pengadaan", monitoringProsesPengadaanImpl.DetailProsesPengadaan)
	monitoring.Put("/proses_pengadaan", monitoringProsesPengadaanImpl.PutProsesPengadaan)

	pengadaan.Get("/status", pengadaanController.IndexStatus)
	pengadaan.Get("/type_pengadaan", pengadaanController.IndexType)

}
