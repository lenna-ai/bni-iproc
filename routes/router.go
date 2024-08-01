package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/injector"

	jwtware "github.com/gofiber/contrib/jwt"
)

func Router(app *fiber.App) {
	var allControllers = injector.InitializeController()

	app.Get("arifin/test", allControllers.LoginController.Arifin)
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
	dashboard.Get("/total_pengadaan", allControllers.DashboardController.TotalPengadaan)
	dashboard.Get("/total_pembayaran", allControllers.DashboardController.TotalPembayaran)
	dashboard.Get("/total_vendor", allControllers.DashboardController.TotalVendor)

	dashboard.Get("/anggaran/:anggaran", allControllers.DashboardController.Anggaran)


	dashboardPengadaanOngoing := dashboard.Group("pengadaanOnGoing")
	dashboardPengadaanOngoing.Get("/kewenangan",allControllers.DashboardController.PengadaanOnGoingKewenangan)
	dashboardPengadaanOngoing.Get("/status",allControllers.DashboardController.PengadaanOnGoingStatus)
	dashboardPengadaanOngoing.Get("/metode",allControllers.DashboardController.PengadaanOnGoingMetode)
	dashboardPengadaanOngoing.Get("/keputusan", allControllers.DashboardController.PengadaanOnGoingKeputusan)

	dashboardPengadaanDone := dashboard.Group("pengadaanDone")
	dashboardPengadaanDone.Get("/kewenangan",allControllers.DashboardController.PengadaanOnDoneKewenangan)
	dashboardPengadaanDone.Get("/status",allControllers.DashboardController.PengadaanOnDoneStatus)
	dashboardPengadaanDone.Get("/metode",allControllers.DashboardController.PengadaanOnDoneMetode)
	dashboardPengadaanDone.Get("/tren_pengadaan/:status",allControllers.DashboardController.PengadaanOnDoneTrenPengadaan)

	informasiRekanan := dashboard.Group("informasiRekanan")
	informasiRekanan.Get("/",allControllers.DashboardController.InformasiRekanan)
	informasiRekanan.Get("/data_informasi_rekanan",allControllers.DashboardController.DataInformasiRekanan)

	pengadaan := app.Group("pengadaan")
	pengadaan.Get("/", allControllers.PengadaanController.IndexPengadaan)
	pengadaan.Get("/filter", allControllers.PengadaanController.FilterPengadaan)
	pengadaan.Get("/sum", allControllers.PengadaanController.SumPengadaan)
	pengadaan.Get("/efisiensi", allControllers.PengadaanController.EfisiensiPengadaan)
	pengadaan.Get("/dynamic/table", allControllers.PengadaanController.DynamicPengadaan)

	pengadaan.Get("/status", allControllers.PengadaanController.IndexStatus)
	pengadaan.Get("/type_pengadaan", allControllers.PengadaanController.IndexType)

	rekanan := app.Group("rekanan")
	rekanan.Get("/:jenis_pengadaan", allControllers.DashboardRekananController.Rekanan)
	rekanan.Get("/:jenis_pengadaan/breakdown/:nama_pt", allControllers.DashboardRekananController.BreakdownRekanan)

	pembayaraan := app.Group("pembayaraan")
	pembayaraan.Get("/", allControllers.PembayaranMonitoringController.IndexPembayaran)
	pembayaraan.Get("rekanan", allControllers.PembayaranMonitoringController.IndexRekananPembayaran)
	pembayaraan.Get("filter", allControllers.PembayaranMonitoringController.FilterPengadaan)

	monitoring := app.Group("monitoring")
	monitoring.Get("/jenis_pengadaan", allControllers.MonitoringProsesPengadaan.JenisPengadaan)
	monitoring.Get("/proses_pengadaan/:jenis_pengadaan", allControllers.MonitoringProsesPengadaan.DetailProsesPengadaan)
	monitoring.Put("/proses_pengadaan", allControllers.MonitoringProsesPengadaan.PutProsesPengadaan)

	pembayaranRutin := monitoring.Group("pembayaranRutin")
	pembayaranRutin.Get("/", allControllers.PembayaranRutinController.DetailPembayaranRutin)
	pembayaranRutin.Put("/", allControllers.PembayaranRutinController.PutPembayaranRutin)

	breakdownPembayaranRutin := pembayaranRutin.Group("breakdown")
	breakdownPembayaranRutin.Get("/", allControllers.PembayaranRutinController.DetailBreakdownPembayaranRutin)
	breakdownPembayaranRutin.Put("/", allControllers.PembayaranRutinController.PutBreakdownPembayaranRutin)

	pembayaranPrestasi := app.Group("pembayaranPrestasi")
	pembayaranPrestasi.Post("/", allControllers.PembayaranPrestasiController.DetailPembayaranPrestasi)
	pembayaranPrestasi.Put("/", allControllers.PembayaranPrestasiController.PutPembayaranPrestasi)

	breakdown := pembayaranPrestasi.Group("breakdown")
	breakdown.Post("/", allControllers.PembayaranPrestasiController.DetailBreakdownPembayaranPrestasi)
	breakdown.Put("/", allControllers.PembayaranPrestasiController.PutBreakdownPembayaranPrestasi)
}
