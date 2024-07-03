package breakdown

type BreakdownPembayaranPrestasi struct {
	TERMIN                    string `json:"TERMIN" validate:"required"`
	STATUS_PEMBAYARAN         string `json:"STATUS_PEMBAYARAN" validate:"required"`
	NILAI_TAGIHAN             string `json:"NILAI_TAGIHAN" validate:"required"`
	STATUS_SYARAT_PEMBAYARAN  string `json:"STATUS_SYARAT_PEMBAYARAN" validate:"required"`
	KETERANGAN_JIKA_TERLAMBAT string `json:"KETERANGAN_JIKA_TERLAMBAT" validate:"required"`
}

func (BreakdownPembayaranPrestasi) TableName() string {
	return "MONITORING_PEMBAYARAN_PRESTASI"
}

type RequestBreakdownPembayaranPrestasi struct {
	JENIS_PENGADAAN string `json:"JENIS_PENGADAAN" validate:"required"`
	NAMA_PENGADAAN  string `json:"NAMA_PENGADAAN" validate:"required"`
	NILAI_PENGADAAN string `json:"NILAI_PENGADAAN"`
}
type RequestPutBreakdownPembayaranPrestasi struct {
	RequestBreakdownPembayaranPrestasi
	TERMIN                    string `json:"TERMIN"`
	STATUS_PEMBAYARAN         string `json:"STATUS_PEMBAYARAN" validate:"required"`
	NILAI_TAGIHAN             string `json:"NILAI_TAGIHAN" validate:"required"`
	STATUS_SYARAT_PEMBAYARAN  string `json:"STATUS_SYARAT_PEMBAYARAN" validate:"required"`
	KETERANGAN_JIKA_TERLAMBAT string `json:"KETERANGAN_JIKA_TERLAMBAT" validate:"required"`
}

func (RequestPutBreakdownPembayaranPrestasi) TableName() string {
	return "MONITORING_PEMBAYARAN_PRESTASI"
}
