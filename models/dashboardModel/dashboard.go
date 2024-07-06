package dashboardmodel

type DashboardRekanan struct {
	VendorId	string
	NamaPekerjaan	string
	NamaVendor	string
	ContractId	string
	NomorSpk	string
	NomorKontrak	string
	NilaiKontrak	string
	JatuhTempoPekerjaan string	
	JenisPengadaan	string
	PengadaanId	string
	StageOfWork	string
	Termin	string
	StatusPayment string	
	NilaiTermin	string
	JatuhTempoPembayaran string	
	TanggalInvoiceDiterima string
}

func (DashboardRekanan) TableName() string {
	return "PEMBAYARAN"
}
