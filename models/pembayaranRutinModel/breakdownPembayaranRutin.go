package pembayaranrutinmodel

type BreakdownPembayaranRutin struct {
	MonitoringPembayaranRutinId string `json:"MONITORING_PEMBAYARAN_RUTIN_ID" validate:"required"`
	JadwalPembayaran            string `json:"JADWAL_PEMBAYARAN" validate:"required"`
	StatusPembayaran            string `json:"STATUS_PEMBAYARAN" validate:"required"`
	NilaiTagihan                string `json:"NILAI_TAGIHAN" validate:"required"`
	CeklisDokumen               string `json:"CEKLIS_DOKUMEN"`
	StatusSyaratPembayaran      string `json:"STATUS_SYARAT_PEMBAYARAN" validate:"required"`
	KeteranganJikaTerlambat     string `json:"KETERANGAN_JIKA_TERLAMBAT" validate:"required"`
}

func (BreakdownPembayaranRutin) TableName() string {
	return "BREAKDOWN_MONITORING_PEMBAYARAN_RUTIN"
}
