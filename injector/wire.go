//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"github.com/lenna-ai/bni-iproc/config"
	"github.com/lenna-ai/bni-iproc/controllers"
	dashboardcontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardController"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	dashboardrekanancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardRekananController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	"github.com/lenna-ai/bni-iproc/controllers/loginController"
	monitoringcontroller "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
	pembayaranprestasicontroller "github.com/lenna-ai/bni-iproc/controllers/pembayaranPrestasiController"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
	dashboardrekananrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRekananRepositories"
	dashboardrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardRepositories"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
	loginrepositories "github.com/lenna-ai/bni-iproc/repositories/loginRepositories"
	monitoringrepositories "github.com/lenna-ai/bni-iproc/repositories/monitoringRepositories"
	pembayaranprestasirepositories "github.com/lenna-ai/bni-iproc/repositories/pembayaranPrestasiRepositories"
	dashboardpembayaranservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
	dashboardrekananservices "github.com/lenna-ai/bni-iproc/services/dashboardRekananServices"
	dashboardservices "github.com/lenna-ai/bni-iproc/services/dashboardServices"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
	loginservices "github.com/lenna-ai/bni-iproc/services/loginServices"
	monitoringservices "github.com/lenna-ai/bni-iproc/services/monitoringServices"
	pembayaranprestasiservices "github.com/lenna-ai/bni-iproc/services/pembayaranPrestasiServices"
	"gorm.io/gorm"
)

func ProvideDB() (*gorm.DB) {
	db := config.DB
	return db
}
var setLoginController = wire.NewSet(
	loginrepositories.NewLoginRepository,
	wire.Bind(new(loginrepositories.LoginRepository), new(*loginrepositories.LoginRepositoryImpl)),
	loginservices.NewLdapLoginService,
	wire.Bind(new(loginservices.LdapLoginService),new(*loginservices.LdapLoginServiceImpl)),
	loginController.NewloginController,
	wire.Bind(new(loginController.LoginController), new(*loginController.LoginControllerImpl)),
)

var setDashboardController = wire.NewSet(
	dashboardrepositories.NewDashboardRepository,
	wire.Bind(new(dashboardrepositories.DashboardRepository),new(*dashboardrepositories.DashboardRepositoryImpl)),
	dashboardservices.NewDashboardService,
	wire.Bind(new(dashboardservices.DashboardService), new(*dashboardservices.DashboardServiceImpl)),
	dashboardcontroller.NewDashboardController,
	wire.Bind(new(dashboardcontroller.DashboardController), new(*dashboardcontroller.DashboardControllerImpl)),
)

var setDashboardRekananController = wire.NewSet(
	dashboardrekananrepositories.NewDashboardRekananRepository,
	wire.Bind(new(dashboardrekananrepositories.DashboardRekananRepository),new(*dashboardrekananrepositories.DashboardRekananRepositoryImpl)),
	dashboardrekananservices.NewDashboardRekananService,
	wire.Bind(new(dashboardrekananservices.DashboardRekananService), new(*dashboardrekananservices.DashboardRekananServiceImpl)),
	dashboardrekanancontroller.NewDashboardRekananController,
	wire.Bind(new(dashboardrekanancontroller.DashboardRekanan), new(*dashboardrekanancontroller.DashboardRekananImpl)),
)

var setPengadaanController = wire.NewSet(
	detailpengadaanrepositories.NewDetailPengadaanRepository,
	wire.Bind(new(detailpengadaanrepositories.PengadaanRepository), new(*detailpengadaanrepositories.PengadaanRepositoryImpl)),
	detailpengadaanservices.NewDetailPengadaanService,
	wire.Bind(new(detailpengadaanservices.PengadaanService), new(*detailpengadaanservices.PengadaanServiceImpl)),
	detailpengadaancontroller.NewDetailPengadaanController,
	wire.Bind(new(detailpengadaancontroller.PengadaanDoneController), new(*detailpengadaancontroller.PengadaanControllerImpl)),
)

var setPembayaranMonitoringController = wire.NewSet(
	dashboardpembayaranrepositories.NewDashboardMonitoringRepository,
	wire.Bind(new(dashboardpembayaranrepositories.PembayaranMonitoringRepository),new(*dashboardpembayaranrepositories.PembayaranMonitoringRepositoryImpl)),
	dashboardpembayaranservices.NewPembayaranMonitoringServices,
	wire.Bind(new(dashboardpembayaranservices.PembayaranMonitoringServices),new(*dashboardpembayaranservices.PembayaranMonitoringServicesImpl)),
	dashboardpembayarancontroller.NewPembayaranMonitoringController,
	wire.Bind(new(dashboardpembayarancontroller.PembayaranMonitoringController), new(*dashboardpembayarancontroller.PembayaranMonitoringControllerImpl)),
)

var setMonitoringProsesPengadaanController = wire.NewSet(
	monitoringrepositories.NewMonitoringProsesPengadaan,
	wire.Bind(new(monitoringrepositories.MonitoringProsesPengadaan), new(*monitoringrepositories.MonitoringProsesPengadaanImpl)),
	monitoringservices.NewMonitoringProsesPengadaan,
	wire.Bind(new(monitoringservices.MonitoringProsesPengadaan), new(*monitoringservices.MonitoringProsesPengadaanImpl)),
	monitoringcontroller.NewMonitoringProsesPengadaan,
	wire.Bind(new(monitoringcontroller.MonitoringProsesPengadaan), new(*monitoringcontroller.MonitoringProsesPengadaanImpl)),
)

var setPembayaranRutinController  = wire.NewSet(
	monitoringrepositories.NewPembayaranRutinRepository,
	wire.Bind(new(monitoringrepositories.PembayaranRutinRepository), new(*monitoringrepositories.PembayaranRutinRepositoryImpl)),
	monitoringservices.NewPembayaranRutinService,
	wire.Bind(new(monitoringservices.PembayaranRutinService), new(*monitoringservices.PembayaranRutinServiceImpl)),
	monitoringcontroller.NewPembayaranRutinController,
	wire.Bind(new(monitoringcontroller.PembayaranRutinController), new(*monitoringcontroller.PembayaranRutinControllerImpl)),
)

var setPembayaranPrestasiController = wire.NewSet(
	pembayaranprestasirepositories.NewPembayaranPrestasiRepository,
	wire.Bind(new(pembayaranprestasirepositories.PembayaranPrestasiRepository), new(*pembayaranprestasirepositories.PembayaranPrestasiRepositoryImpl)),
	pembayaranprestasiservices.NewPembayaranPrestasiService,
	wire.Bind(new(pembayaranprestasiservices.PembayaranPrestasiService), new(*pembayaranprestasiservices.PembayaranPrestasiServiceImpl)),
	pembayaranprestasicontroller.NewPembayaranPrestasiController,
	wire.Bind(new(pembayaranprestasicontroller.PembayaranPrestasiController),new(*pembayaranprestasicontroller.PembayaranPrestasiControllerImpl)),
)

var setAllControllers = wire.NewSet(
	ProvideDB,
	setLoginController,
	setDashboardController,
	setDashboardRekananController,
	setPengadaanController,
	setPembayaranMonitoringController,
	setMonitoringProsesPengadaanController,
	setPembayaranRutinController,
	setPembayaranPrestasiController,
	wire.Struct(new(controllers.AllControllers), "*"),
)

func InitializeController() (*controllers.AllControllers){
	// wire.Build(setLoginController, controllers.NewAllControllers)
	wire.Build(setAllControllers)
	return &controllers.AllControllers{}
}

// var setDashboardController = wire.NewSet(
// 	// wire.NewSet(ProvideDB, dashboardrepositories.NewDashboardRepository),
// 	dashboardrepositories.NewDashboardRepository,
// 	wire.Bind(new(dashboardrepositories.DashboardRepository),new(*dashboardrepositories.DashboardRepositoryImpl)),
// 	dashboardservices.NewDashboardService,
// 	wire.Bind(new(dashboardservices.DashboardService), new(*dashboardservices.DashboardServiceImpl)),
// 	dashboardcontroller.NewDashboardController,
// 	wire.Bind(new(dashboardcontroller.DashboardController), new(*dashboardcontroller.DashboardControllerImpl)),
// 	wire.Struct(new(controllers.AllControllers), "DashboardController"),
// )
