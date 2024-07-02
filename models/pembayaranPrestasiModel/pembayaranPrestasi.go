package pembayaranprestasimodel

type PembayaranPrestasi struct {
	ID                    int
	NAMA                  string
	NILAI_PENGADAAN_HASIL string
	PEMBAYARAN_TERAKHIR   string
	STATUS_JATUH_TEMPO    string
}

func (PembayaranPrestasi) TableName() string {
	return "MONITORING_PEMBAYARAN_RUTIN"
}
