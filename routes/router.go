package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	appconfig "github.com/lenna-ai/bni-iproc/config/appConfig"
	"github.com/lenna-ai/bni-iproc/controllers"

	jwtware "github.com/gofiber/contrib/jwt"
)

func Router(app *fiber.App) {
	var allControllers = new(controllers.AllControllers)
	appconfig.InitControllerServiceRepository(allControllers)

	login := app.Group("login")
	login.Post("ldap",allControllers.LoginController.Ldap)
	login.Post("vendor",allControllers.LoginController.Vendor)

	// comments JWT 
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_TOKEN"))},
		ErrorHandler: allControllers.LoginController.ErrorHandler,
	}))

	app.Get("/me", allControllers.LoginController.MeJwt)

	dashboard := app.Group("dashboard")
	dashboard.Get("/total_pengadaan", allControllers.DashboardControllerImpl.TotalPengadaan)
	dashboard.Get("/total_pembayaran", allControllers.DashboardControllerImpl.TotalPembayaran)
	dashboard.Get("/total_vendor", allControllers.DashboardControllerImpl.TotalVendor)

	dashboardPengadaanOngoing := dashboard.Group("pengadaanOnGoing")
	dashboardPengadaanOngoing.Get("/kewenangan",allControllers.DashboardControllerImpl.PengadaanOnGoingKewenangan)
	dashboardPengadaanOngoing.Get("/status",allControllers.DashboardControllerImpl.PengadaanOnGoingStatus)
	dashboardPengadaanOngoing.Get("/metode",allControllers.DashboardControllerImpl.PengadaanOnGoingMetode)
	dashboardPengadaanOngoing.Get("/keputusan", allControllers.DashboardControllerImpl.PengadaanOnGoingKeputusan)

	dashboardPengadaanDone := dashboard.Group("pengadaanDone")
	dashboardPengadaanDone.Get("/kewenangan",allControllers.DashboardControllerImpl.PengadaanOnDoneKewenangan)
	dashboardPengadaanDone.Get("/status",allControllers.DashboardControllerImpl.PengadaanOnDoneStatus)
	dashboardPengadaanDone.Get("/metode",allControllers.DashboardControllerImpl.PengadaanOnDoneMetode)

	informasiRekanan := dashboard.Group("informasiRekanan")
	informasiRekanan.Get("/",allControllers.DashboardControllerImpl.InformasiRekanan)
	informasiRekanan.Get("/data_informasi_rekanan",allControllers.DashboardControllerImpl.DataInformasiRekanan)

	pengadaan := app.Group("pengadaan")
	pengadaan.Get("/", allControllers.PengadaanControllerImpl.IndexPengadaan)
	pengadaan.Get("/filter", allControllers.PengadaanControllerImpl.FilterPengadaan)
	pengadaan.Get("/sum", allControllers.PengadaanControllerImpl.SumPengadaan)
	pengadaan.Get("/efisiensi", allControllers.PengadaanControllerImpl.EfisiensiPengadaan)

	pengadaan.Get("/status", allControllers.PengadaanControllerImpl.IndexStatus)
	pengadaan.Get("/type_pengadaan", allControllers.PengadaanControllerImpl.IndexType)

	rekanan := app.Group("rekanan")
	rekanan.Get("/:jenis_pengadaan", allControllers.DashboardRekananController.Rekanan)
	rekanan.Get("/breakdown/:nama_pt", allControllers.DashboardRekananController.BreakdownRekanan)

	pembayaraan := app.Group("pembayaraan")
	pembayaraan.Get("/", allControllers.PembayaranMonitoringControllerImpl.IndexPembayaran)
	pembayaraan.Get("rekanan", allControllers.PembayaranMonitoringControllerImpl.IndexRekananPembayaran)
	pembayaraan.Get("filter", allControllers.PembayaranMonitoringControllerImpl.FilterPengadaan)

	monitoring := app.Group("monitoring")
	monitoring.Get("/jenis_pengadaan", allControllers.MonitoringProsesPengadaanImpl.JenisPengadaan)
	monitoring.Get("/proses_pengadaan", allControllers.MonitoringProsesPengadaanImpl.DetailProsesPengadaan)
	monitoring.Put("/proses_pengadaan", allControllers.MonitoringProsesPengadaanImpl.PutProsesPengadaan)

	pembayaranRutin := monitoring.Group("pembayaranRutin")
	pembayaranRutin.Get("/", allControllers.PembayaranRutinControllerImpl.DetailPembayaranRutin)
	pembayaranRutin.Put("/", allControllers.PembayaranRutinControllerImpl.PutPembayaranRutin)

	breakdownPembayaranRutin := pembayaranRutin.Group("breakdown")
	breakdownPembayaranRutin.Get("/", allControllers.PembayaranRutinControllerImpl.DetailBreakdownPembayaranRutin)
	breakdownPembayaranRutin.Put("/", allControllers.PembayaranRutinControllerImpl.PutBreakdownPembayaranRutin)

	pembayaranPrestasi := app.Group("pembayaranPrestasi")
	pembayaranPrestasi.Get("/", allControllers.PembayaranPrestasiControllerImpl.DetailPembayaranPrestasi)
	pembayaranPrestasi.Put("/", allControllers.PembayaranPrestasiControllerImpl.PutPembayaranPrestasi)

	breakdown := pembayaranPrestasi.Group("breakdown")
	breakdown.Get("/", allControllers.PembayaranPrestasiControllerImpl.DetailBreakdownPembayaranPrestasi)
	breakdown.Put("/", allControllers.PembayaranPrestasiControllerImpl.PutBreakdownPembayaranPrestasi)
}
