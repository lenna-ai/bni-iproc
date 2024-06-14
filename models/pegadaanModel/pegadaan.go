package detailmodel

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

type DataResultSumPengadaan struct {
	SUM_NILAI_PENGADAAN_HASIL string `gorm:"column:NILAI_PENGADAAN_HASIL"`
	GROUP_PENGADAAN           string `gorm:"column:JENIS_PENGADAAN"`
}

func (Pengadaan) TableName() string {
	return "PENGADAAN"
}
