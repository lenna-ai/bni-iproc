package pembayaranprestasirepositories

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel/breakdown"
)

func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi, requestPembayaranPrestasi *pembayaranprestasimodel.RequestPembayaranPrestasi) error {
	if err := pembayaranPrestasiRepositoryImpl.DB.Where("JENIS_PENGADAAN = ?", requestPembayaranPrestasi.JENIS_PENGADAAN).Find(pembayaranPrestasi).Error; err != nil {
		fmt.Println("pembayaranPrestasiRepositoryImpl.DB.Find(pembayaranPrestasi).Error")
		return err
	}
	return nil
}
func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) PutPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *pembayaranprestasimodel.PembayaranPrestasi) error {
	var whereQuery string
	fmt.Println(whereQuery)
	if pembayaranPrestasi.NILAI_PENGADAAN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN IS NULL", pembayaranPrestasi.NAMA_PENGADAAN, pembayaranPrestasi.JENIS_PENGADAAN)
	} else {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN = '%s'", pembayaranPrestasi.NAMA_PENGADAAN, pembayaranPrestasi.JENIS_PENGADAAN, pembayaranPrestasi.NILAI_PENGADAAN)
	}
	updateProsesPengadaanModel := pembayaranPrestasiRepositoryImpl.DB.Where(whereQuery).Updates(pembayaranPrestasi)
	fmt.Println(updateProsesPengadaanModel.RowsAffected)
	if updateProsesPengadaanModel.RowsAffected < 1 {
		return errors.New("Data not found")
	}
	if err := updateProsesPengadaanModel.Error; err != nil {
		return err
	}
	return nil
}

func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) DetailBreakdownPembayaranPrestasi(c *fiber.Ctx, breakdownPembayaraanPrestasi *[]breakdown.BreakdownPembayaranPrestasi, breakdownRequestBreakdownPembayaranPrestasi *breakdown.RequestBreakdownPembayaranPrestasi) error {
	var whereQuery string
	if breakdownRequestBreakdownPembayaranPrestasi.NILAI_PENGADAAN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN IS NULL", breakdownRequestBreakdownPembayaranPrestasi.NAMA_PENGADAAN, breakdownRequestBreakdownPembayaranPrestasi.JENIS_PENGADAAN)
	} else {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN = '%s'", breakdownRequestBreakdownPembayaranPrestasi.NAMA_PENGADAAN, breakdownRequestBreakdownPembayaranPrestasi.JENIS_PENGADAAN, breakdownRequestBreakdownPembayaranPrestasi.NILAI_PENGADAAN)
	}
	if err := pembayaranPrestasiRepositoryImpl.DB.Where(whereQuery).Find(breakdownPembayaraanPrestasi).Error; err != nil {
		fmt.Println("pembayaranPrestasiRepositoryImpl.DB.Find(breakdownPembayaraanPrestasi).Error")
		return err
	}
	return nil
}
func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) PutBreakdownPembayaranPrestasi(c *fiber.Ctx) error {
	fmt.Println("REPO PutBreakdownPembayaranPrestasi")
	return nil
}
