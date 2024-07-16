package formatters

type PutPengadaanFormatter struct {
	Nama                    string `json:"Nama" validate:"required"`
	Metode              string `json:"Metode" validate:"required"`
	Tahapan             string `json:"Tahapan" validate:"required"`
	Status              string `json:"Status" validate:"required"`
	ScheduleEndDate   string `json:"ScheduleEndDate" validate:"required"`
	ScheduleStartDate string `json:"ScheduleStartDate" validate:"required"`
	KeteranganJikaTerlambat string `json:"Keterangan_jika_terlambat" validate:"required,lt=400"`
}

func (PutPengadaanFormatter) TableName() string {
	return "MONITORING_PROSES_PENGADAAN"
}
