package pembayaranprestasirepositories

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
)

func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi) error {
	if err := pembayaranPrestasiRepositoryImpl.DB.Find(pembayaranPrestasi).Error; err != nil {
		fmt.Println("pembayaranPrestasiRepositoryImpl.DB.Find(pembayaranPrestasi).Error")
		return err
	}
	return nil
}
func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) PutPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *pembayaranprestasimodel.PembayaranPrestasi) error {
	var whereQuery string
	if pembayaranPrestasi.NILAI_PENGADAAN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and NILAI_PENGADAAN IS NULL", pembayaranPrestasi.NAMA_PENGADAAN)
	} else {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and NILAI_PENGADAAN = '%s'", pembayaranPrestasi.NAMA_PENGADAAN, pembayaranPrestasi.NILAI_PENGADAAN)
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
