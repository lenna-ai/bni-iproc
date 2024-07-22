package formatters

import "time"

type PutPengadaanFormatter struct {
	ID string `json:"ID"`
	PROCUREMENT_ID string `json:"PROCUREMENT_ID" validate:"required"`
	STATUS string `json:"STATUS" validate:"required"`
	STATUS_PENGADAAN_PROMOTS string `json:"STATUS_PENGADAAN_PROMOTS" validate:"required"`
	KETERANGAN_JIKA_TERLAMBAT string `json:"KETERANGAN_JIKA_TERLAMBAT" validate:"required,lt=400"`
	DELETED_AT *time.Time `json:"DELETED_AT"`
}


func (PutPengadaanFormatter) TableName() string {
	return "MONITORING_PROSES_PENGADAAN_NEW"
}
