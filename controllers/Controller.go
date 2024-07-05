package controllers

import (
	dashboardcontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardController"
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	monitoringcontroller "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
	pembayaranprestasicontroller "github.com/lenna-ai/bni-iproc/controllers/pembayaranPrestasiController"
)

type AllControllers struct {
	DashboardControllerImpl            *dashboardcontroller.DashboardControllerImpl
	PengadaanControllerImpl            *detailpengadaancontroller.PengadaanControllerImpl
	PembayaranMonitoringControllerImpl *dashboardpembayarancontroller.PembayaranMonitoringControllerImpl
	MonitoringProsesPengadaanImpl      *monitoringcontroller.MonitoringProsesPengadaanImpl
	PembayaranRutinControllerImpl      *monitoringcontroller.PembayaranRutinControllerImpl
	PembayaranPrestasiControllerImpl   *pembayaranprestasicontroller.PembayaranPrestasiControllerImpl
}
