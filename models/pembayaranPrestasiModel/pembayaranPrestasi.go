package pembayaranprestasimodel

type PembayaranPrestasi struct {
	ID                         int
	JENIS_PENGADAAN            string `json:"JENIS_PENGADAAN" validate:"required"`
	NAMA_PENGADAAN             string `json:"NAMA_PENGADAAN" validate:"required"`
	NILAI_PENGADAAN            string `json:"NILAI_PENGADAAN" validate:"required"`
	STATUS_PEMBAYARAN_TERAKHIR string `json:"STATUS_PEMBAYARAN_TERAKHIR" validate:"required"`
	STATUS_JATUH_TEMPO         string `json:"STATUS_JATUH_TEMPO" validate:"required"`
}

func (PembayaranPrestasi) TableName() string {
	return "MONITORING_PEMBAYARAN_PRESTASI"
}
