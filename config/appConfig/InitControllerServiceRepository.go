package appconfig

import (
	"github.com/lenna-ai/bni-iproc/config"
	"github.com/lenna-ai/bni-iproc/controllers"
	dashboardcontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardController"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	monitoringController "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
	pembayaranprestasicontroller "github.com/lenna-ai/bni-iproc/controllers/pembayaranPrestasiController"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
	dashboardrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRepositories"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
	monitoringrepositories "github.com/lenna-ai/bni-iproc/repositories/monitoringRepositories"
	pembayaranprestasirepositories "github.com/lenna-ai/bni-iproc/repositories/pembayaranPrestasiRepositories"
	dashboardpembayaranservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
	dashboardservices "github.com/lenna-ai/bni-iproc/services/dashboardServices"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
	monitoringService "github.com/lenna-ai/bni-iproc/services/monitoringServices"
	pembayaranprestasiservices "github.com/lenna-ai/bni-iproc/services/pembayaranPrestasiServices"
)

func InitControllerServiceRepository(allControllers *controllers.AllControllers) {
	db := config.DB

	dashboardRepository := dashboardrepositories.NewDashboardRepository(db)
	dashboardservices := dashboardservices.NewDashboardService(dashboardRepository)
	dashboardcontroller := dashboardcontroller.NewDashboardController(dashboardservices)

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

	allControllers.DashboardControllerImpl = dashboardcontroller
	allControllers.MonitoringProsesPengadaanImpl = monitoringProsesPengadaanController
	allControllers.PembayaranMonitoringControllerImpl = dashboardMonitoringController
	allControllers.PembayaranRutinControllerImpl = PembayaranRutinController
	allControllers.PengadaanControllerImpl = detailPengadaanFilterController
	allControllers.PembayaranPrestasiControllerImpl = pembayaranPrestasiController

	// cannot use pembayaranMonitoringServices (variable of type *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl) as dashboardpembayaraanservices.PembayaranMonitoringServices value in argument to dashboardpembayaraancontroller.NewPembayaranMonitoringController: *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl does not implement dashboardpembayaraanservices.PembayaranMonitoringServices (missing method IndexPengadaan)
	// return allControllers
}
