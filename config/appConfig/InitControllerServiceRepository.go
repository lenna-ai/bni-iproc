package appconfig

import (
	"github.com/lenna-ai/bni-iproc/config"
	"github.com/lenna-ai/bni-iproc/controllers"
	dashboardcontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardController"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	dashboardrekanancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardRekananController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	"github.com/lenna-ai/bni-iproc/controllers/loginController"
	monitoringController "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
	pembayaranprestasicontroller "github.com/lenna-ai/bni-iproc/controllers/pembayaranPrestasiController"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
	dashboardrekananrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRekananRepositories"
	dashboardrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRepositories"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
	monitoringrepositories "github.com/lenna-ai/bni-iproc/repositories/monitoringRepositories"
	pembayaranprestasirepositories "github.com/lenna-ai/bni-iproc/repositories/pembayaranPrestasiRepositories"
	dashboardpembayaranservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
	dashboardrekananservices "github.com/lenna-ai/bni-iproc/services/dashboardRekananServices"
	dashboardservices "github.com/lenna-ai/bni-iproc/services/dashboardServices"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
	loginservices "github.com/lenna-ai/bni-iproc/services/loginServices"
	monitoringService "github.com/lenna-ai/bni-iproc/services/monitoringServices"
	pembayaranprestasiservices "github.com/lenna-ai/bni-iproc/services/pembayaranPrestasiServices"
)

func InitControllerServiceRepository(allControllers *controllers.AllControllers) {
	db := config.DB

	loginServices := loginservices.NewLdapLoginService()
	loginController := loginController.NewloginController(loginServices)

	dashboardRepository := dashboardrepositories.NewDashboardRepository(db)
	dashboardservices := dashboardservices.NewDashboardService(dashboardRepository)
	dashboardcontroller := dashboardcontroller.NewDashboardController(dashboardservices)

	dashboardRekananRepository :=dashboardrekananrepositories.NewDashboardRekananRepository(db)
	dashboardRekananService := dashboardrekananservices.NewDashboardRekananService(dashboardRekananRepository)
	dashboardRekanan := dashboardrekanancontroller.NewDashboardRekananController(dashboardRekananService)

	detailPengadaanFilterRepository := detailpengadaanrepositories.NewDetailPengadaanRepository(db)
	detailPengadaanFilterService := detailpengadaanservices.NewDetailPengadaanService(detailPengadaanFilterRepository)
	detailPengadaanFilterController := detailpengadaancontroller.NewDetailPengadaanController(detailPengadaanFilterService)

	pembayaranMonitoringRepository := dashboardpembayaranrepositories.NewDashboardMonitoringRepository(db)
	pembayaranMonitoringServices := dashboardpembayaranservices.NewPembayaranMonitoringServices(pembayaranMonitoringRepository)
	dashboardMonitoringController := dashboardpembayarancontroller.NewPembayaranMonitoringController(pembayaranMonitoringServices)

	monitoringProsesPengadaanRepository := monitoringrepositories.NewMonitoringProsesPengadaan(db)
	monitoringProsesPengadaanService := monitoringService.NewMonitoringProsesPengadaan(monitoringProsesPengadaanRepository)
	monitoringProsesPengadaanController := monitoringController.NewMonitoringProsesPengadaan(monitoringProsesPengadaanService)

	monitoringPembayaranRutinRepository := monitoringrepositories.NewPembayaranRutinRepository(db)
	monitoringPembayaranRutinService := monitoringService.NewPembayaranRutinService(monitoringPembayaranRutinRepository)
	PembayaranRutinController := monitoringController.NewPembayaranRutinController(monitoringPembayaranRutinService)

	pembayaranPrestasiRepository := pembayaranprestasirepositories.NewPembayaranPrestasiRepository(db)
	pembayaranPrestasiServices := pembayaranprestasiservices.NewPembayaranPrestasiService(pembayaranPrestasiRepository)
	pembayaranPrestasiController := pembayaranprestasicontroller.NewPembayaranPrestasiController(pembayaranPrestasiServices)

	allControllers.LoginController = loginController
	allControllers.DashboardControllerImpl = dashboardcontroller
	allControllers.DashboardRekananController = dashboardRekanan
	allControllers.MonitoringProsesPengadaanImpl = monitoringProsesPengadaanController
	allControllers.PembayaranMonitoringControllerImpl = dashboardMonitoringController
	allControllers.PembayaranRutinControllerImpl = PembayaranRutinController
	allControllers.PengadaanControllerImpl = detailPengadaanFilterController
	allControllers.PembayaranPrestasiControllerImpl = pembayaranPrestasiController

	// cannot use pembayaranMonitoringServices (variable of type *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl) as dashboardpembayaraanservices.PembayaranMonitoringServices value in argument to dashboardpembayaraancontroller.NewPembayaranMonitoringController: *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl does not implement dashboardpembayaraanservices.PembayaranMonitoringServices (missing method IndexPengadaan)
	// return allControllers
}
