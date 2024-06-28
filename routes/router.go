package routes

import (
	"github.com/gofiber/fiber/v2"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
	"github.com/lenna-ai/bni-iproc/controllers"
)

func Router(app *fiber.App) {
	var allControllers = new(controllers.AllControllers)
	appconfig.InitControllerServiceRepository(allControllers)
	pengadaan := app.Group("pengadaan")
	pengadaan.Get("/", allControllers.PengadaanControllerImpl.IndexPengadaan)
	pengadaan.Get("/filter", allControllers.PengadaanControllerImpl.FilterPengadaan)
	pengadaan.Get("/sum", allControllers.PengadaanControllerImpl.SumPengadaan)
	pengadaan.Get("/efisiensi", allControllers.PengadaanControllerImpl.EfisiensiPengadaan)

	pembayaraan := app.Group("pembayaraan")
	pembayaraan.Get("/", allControllers.PembayaranMonitoringControllerImpl.IndexPembayaran)
	pembayaraan.Get("rekanan", allControllers.PembayaranMonitoringControllerImpl.IndexRekananPembayaran)
	pembayaraan.Get("filter", allControllers.PembayaranMonitoringControllerImpl.FilterPengadaan)

	monitoring := app.Group("monitoring")
	monitoring.Get("/jenis_pengadaan", allControllers.MonitoringProsesPengadaanImpl.JenisPengadaan)
	monitoring.Get("/proses_pengadaan", allControllers.MonitoringProsesPengadaanImpl.DetailProsesPengadaan)
	monitoring.Put("/proses_pengadaan", allControllers.MonitoringProsesPengadaanImpl.PutProsesPengadaan)

	pengadaan.Get("/status", allControllers.PengadaanControllerImpl.IndexStatus)
	pengadaan.Get("/type_pengadaan", allControllers.PengadaanControllerImpl.IndexType)

	pembayaranRutin := monitoring.Group("pembayaranRutin")
	pembayaranRutin.Get("/", allControllers.PembayaranRutinControllerImpl.DetailPembayaranRutin)
	pembayaranRutin.Put("/", allControllers.PembayaranRutinControllerImpl.PutPembayaranRutin)

}
