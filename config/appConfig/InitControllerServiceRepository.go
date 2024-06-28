package appconfig

import (
	"github.com/lenna-ai/bni-iproc/config"
	"github.com/lenna-ai/bni-iproc/controllers"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	monitoringController "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
	monitoringrepositories "github.com/lenna-ai/bni-iproc/repositories/monitoringRepositories"
	dashboardpembayaranservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
	monitoringService "github.com/lenna-ai/bni-iproc/services/monitoringServices"
)

func InitControllerServiceRepository(allControllers *controllers.AllControllers) {
	db := config.DB
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

	allControllers.MonitoringProsesPengadaanImpl = monitoringProsesPengadaanController
	allControllers.PembayaranMonitoringControllerImpl = dashboardMonitoringController
	allControllers.PembayaranRutinControllerImpl = PembayaranRutinController
	allControllers.PengadaanControllerImpl = detailPengadaanFilterController

	// cannot use pembayaranMonitoringServices (variable of type *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl) as dashboardpembayaraanservices.PembayaranMonitoringServices value in argument to dashboardpembayaraancontroller.NewPembayaranMonitoringController: *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl does not implement dashboardpembayaraanservices.PembayaranMonitoringServices (missing method IndexPengadaan)
	// return allControllers
}
