package pegadaanmodel

import "github.com/lenna-ai/bni-iproc/models/prosesPengadaanModel/formatters"

type Pengadaan struct {
	Procurement_id          int
	Jenis_pengadaan         string
	Jenis_pengadaan_id      string
	Nama                    string
	PIC_Pengadaan           string
	Unit_Pengadaan          string
	Metode                  string
	Tahapan                 string
	Status                  string
	Schedule_start_date     string
	Schedule_end_date       string
	Status_pengadaan        string
	Status_pengadaan_id     string
	Check_dokumen           string
	Nilai_pengadaan_inisasi string
	Nilai_pengadaan_hasil   string
	Stage_of_work           string
	Termin                  string
	Kewenangan_pengadaan    string
	Nomor_spk               string
	Nama_vendor             string
	Nilai_spk               string
	Mata_anggaran           string
	Sla_in_days             string
}

type PengadaanFilter struct {
	PROCUREMENT_ID int
	JENIS_PENGADAAN string
	JENIS_PENGADAAN_ID string
	NAMA string
	PIC_PENGADAAN string
	PIC_PENGADAAN_ID string
	UNIT_PENGADAAN string
	METODE string
	TAHAPAN string
	STATUS string
	SCHEDULE_START_DATE string
	SCHEDULE_END_DATE string
	STATUS_PENGADAAN string
	STATUS_PENGADAAN_ID int
	CHECK_DOKUMEN string
	NILAI_PENGADAAN_INISASI string
	NILAI_PENGADAAN_HASIL string
	KEWENANGAN_PENGADAAN string
	NOMOR_SPK string
	NILAI_SPK string
	NAMA_VENDOR string
	MATA_ANGGARAN string
	ITEM_NAME string
	POST_DATE_SPK string
	LETTER_DATE_SPK string
	CREATED_AT string
	SLA_IN_DAYS string
	MonitoringProses  []formatters.PutPengadaanFormatter `gorm:"foreignKey:PROCUREMENT_ID;references:PROCUREMENT_ID"`
}

type DataResultSumPengadaan struct {
	ESTIMASI_NILAI_PENGADAAN string `gorm:"column:ESTIMASI_NILAI_PENGADAAN"`
	NILAI_SPK                string `gorm:"column:NILAI_SPK"`
	GROUP_PENGADAAN          string `gorm:"column:JENIS_PENGADAAN"`
}

func (Pengadaan) TableName() string {
	return "PENGADAAN"
}

func (PengadaanFilter) TableName() string {
	return "PENGADAAN_FILTER"
}
