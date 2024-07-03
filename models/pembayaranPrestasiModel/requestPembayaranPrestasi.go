package pembayaranprestasimodel

type RequestPembayaranPrestasi struct {
	JENIS_PENGADAAN string `json:"JENIS_PENGADAAN" validate:"required"`
}
