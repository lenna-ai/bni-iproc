package pembayaranrutinmodel

type PembayaranRutin struct {
	Nama                string `json:"nama" validate:"required"`
	NilaiPengadaanHasil string `json:"nilai_pengadaan_hasil" validate:"required"`
	PembayaranTerakhir  string `json:"pembayaran_terakhir" validate:"required"`
	StatusJatuhTempo    string `json:"status_jatuh_tempo" validate:"required"`
	NominalPembayaran   string `json:"nominal_pembayaran" validate:"required"`
}

func (PembayaranRutin) TableName() string {
	return "MONITORING_PEMBAYARAN_RUTIN"
}
