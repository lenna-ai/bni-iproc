package formatters

import "time"

type PutPengadaanFormatter struct {
	ID string `json:"ID"`
	PROCUREMENT_ID string `json:"PROCUREMENT_ID" validate:"required" gorm:"index" column:"PROCUREMENT_ID"`
	STATUS string `json:"STATUS" validate:"required" column:"STATUS"`
	STATUS_PENGADAAN_PROMOTS string `json:"STATUS_PENGADAAN_PROMOTS" validate:"required" column:"PROCUREMENT_ID"`
	KETERANGAN_JIKA_TERLAMBAT string `json:"KETERANGAN_JIKA_TERLAMBAT" validate:"required,lt=400" column:"KETERANGAN_JIKA_TERLAMBAT"`
	DELETED_AT *time.Time `json:"DELETED_AT"`
}


func (PutPengadaanFormatter) TableName() string {
	return "MONITORING_PROSES_PENGADAAN_NEW"
}

