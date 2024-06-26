package formatters

type PutPengadaanFormatter struct {
	Nama                    string `json:"Nama" validate:"required"`
	KeteranganJikaTerlambat string `json:"Keterangan_jika_terlambat" validate:"required,lt=400"`
}

func (PutPengadaanFormatter) TableName() string {
	return "MONITORING_PROSES_PENGADAAN"
}
