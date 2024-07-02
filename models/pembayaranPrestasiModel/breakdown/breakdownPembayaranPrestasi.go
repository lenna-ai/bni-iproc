package breakdown

type BreakdownPembayaranPrestasi struct {
	TERMIN                    string `json:"TERMIN" validate:"required"`
	STATUS_PEMBAYARAN         string `json:"STATUS_PEMBAYARAN" validate:"required"`
	NILAI_TAGIHAN             string `json:"NILAI_TAGIHAN" validate:"required"`
	STATUS_SYARAT_PEMBAYARAN  string `json:"STATUS_SYARAT_PEMBAYARAN" validate:"required"`
	KETERANGAN_JIKA_TERLAMBAT string `json:"KETERANGAN_JIKA_TERLAMBAT" validate:"required"`
}
