package controllers

import (
	dashboardpembayarancontroller "github.com/lenna-ai/bni-iproc/controllers/dashboardPembayaraanController"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	monitoringcontroller "github.com/lenna-ai/bni-iproc/controllers/monitoringController"
)

type AllControllers struct {
	PengadaanControllerImpl            *detailpengadaancontroller.PengadaanControllerImpl
	PembayaranMonitoringControllerImpl *dashboardpembayarancontroller.PembayaranMonitoringControllerImpl
	MonitoringProsesPengadaanImpl      *monitoringcontroller.MonitoringProsesPengadaanImpl
	PembayaranRutinControllerImpl      *monitoringcontroller.PembayaranRutinControllerImpl
}
