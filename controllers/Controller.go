package controllers

import (
	dashboardcontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardController"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	dashboardrekanancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardRekananController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	"github.com/lenna-ai/bni-iproc/controllers/loginController"
	monitoringcontroller "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
	pembayaranprestasicontroller "github.com/lenna-ai/bni-iproc/controllers/pembayaranPrestasiController"
)

type AllControllers struct {
	LoginController 				*loginController.LoginControllerImpl
	DashboardController            	*dashboardcontroller.DashboardControllerImpl
	DashboardRekananController 		*dashboardrekanancontroller.DashboardRekananImpl
	PengadaanController            	*detailpengadaancontroller.PengadaanControllerImpl
	PembayaranMonitoringController 	*dashboardpembayarancontroller.PembayaranMonitoringControllerImpl
	MonitoringProsesPengadaan      	*monitoringcontroller.MonitoringProsesPengadaanImpl
	PembayaranRutinController      	*monitoringcontroller.PembayaranRutinControllerImpl
	PembayaranPrestasiController   	*pembayaranprestasicontroller.PembayaranPrestasiControllerImpl
}

// func NewTestLoginController(
// 	loginController *loginController.LoginControllerImpl,
// 	dashboardController *dashboardcontroller.DashboardControllerImpl,
// 	dashboardRekananController 		   *dashboardrekanancontroller.DashboardRekananImpl,
// 	pengadaanController            *detailpengadaancontroller.PengadaanControllerImpl,
// 	pembayaranMonitoringController *dashboardpembayarancontroller.PembayaranMonitoringControllerImpl,
// 	monitoringProsesPengadaan      *monitoringcontroller.MonitoringProsesPengadaanImpl,
// 	pembayaranRutinController     *monitoringcontroller.PembayaranRutinControllerImpl,
// 	) *DefineController {
//     return &DefineController{
//         LoginController: loginController,
//         DashboardController: dashboardController,
//         DashboardRekananController: dashboardRekananController,
//         PengadaanController: pengadaanController,
//         PembayaranMonitoringController: pembayaranMonitoringController,
// 		MonitoringProsesPengadaan: monitoringProsesPengadaan,
// 		PembayaranRutinController: pembayaranRutinController,
//     }
// }



// type AllControllers struct {
// 	LoginController *loginController.LoginControllerImpl
// 	DashboardControllerImpl            *dashboardcontroller.DashboardControllerImpl
// 	DashboardRekananController 		   *dashboardrekanancontroller.DashboardRekananImpl
// 	PengadaanControllerImpl            *detailpengadaancontroller.PengadaanControllerImpl
// 	PembayaranMonitoringControllerImpl *dashboardpembayarancontroller.PembayaranMonitoringControllerImpl
// 	MonitoringProsesPengadaanImpl      *monitoringcontroller.MonitoringProsesPengadaanImpl
// 	PembayaranRutinControllerImpl      *monitoringcontroller.PembayaranRutinControllerImpl
// 	PembayaranPrestasiControllerImpl   *pembayaranprestasicontroller.PembayaranPrestasiControllerImpl
// }
