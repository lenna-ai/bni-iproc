package appconfig

import (
	"github.com/lenna-ai/bni-iproc/config"
	dashboardpembayaraancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
	dashboardpembayaranservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
)

func InitControllerServiceRepository() (*detailpengadaancontroller.PengadaanControllerImpl, *dashboardpembayaraancontroller.PembayaranMonitoringControllerImpl) {
	db := config.DB
	detailPengadaanFilterRepository := detailpengadaanrepositories.NewDetailPengadaanRepository(db)
	detailPengadaanFilterService := detailpengadaanservices.NewDetailPengadaanService(detailPengadaanFilterRepository)
	detailPengadaanFilterController := detailpengadaancontroller.NewDetailPengadaanController(detailPengadaanFilterService)

	pembayaranMonitoringRepository := dashboardpembayaranrepositories.NewDashboardMonitoringRepository(db)
	pembayaranMonitoringServices := dashboardpembayaranservices.NewPembayaranMonitoringServices(pembayaranMonitoringRepository)
	dashboardMonitoringController := dashboardpembayarancontroller.NewPembayaranMonitoringController(pembayaranMonitoringServices)

	// cannot use pembayaranMonitoringServices (variable of type *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl) as dashboardpembayaraanservices.PembayaranMonitoringServices value in argument to dashboardpembayaraancontroller.NewPembayaranMonitoringController: *dashboardpembayaraanservices.PembayaranMonitoringServicesImpl does not implement dashboardpembayaraanservices.PembayaranMonitoringServices (missing method IndexPengadaan)
	return detailPengadaanFilterController, dashboardMonitoringController
}
