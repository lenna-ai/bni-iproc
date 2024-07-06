package controllers

import (
	logincontroller "github.com/lenna-ai/bni-iproc/controllers/LoginController"
	dashboardcontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardController"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	dashboardrekanancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardRekananController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	monitoringcontroller "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
	pembayaranprestasicontroller "github.com/lenna-ai/bni-iproc/controllers/pembayaranPrestasiController"
)

type AllControllers struct {
	LoginControllerImpl            	   *logincontroller.LoginControllerImpl
	DashboardControllerImpl            *dashboardcontroller.DashboardControllerImpl
	DashboardRekananController 		   *dashboardrekanancontroller.DashboardRekananImpl
	PengadaanControllerImpl            *detailpengadaancontroller.PengadaanControllerImpl
	PembayaranMonitoringControllerImpl *dashboardpembayarancontroller.PembayaranMonitoringControllerImpl
	MonitoringProsesPengadaanImpl      *monitoringcontroller.MonitoringProsesPengadaanImpl
	PembayaranRutinControllerImpl      *monitoringcontroller.PembayaranRutinControllerImpl
	PembayaranPrestasiControllerImpl   *pembayaranprestasicontroller.PembayaranPrestasiControllerImpl
}
