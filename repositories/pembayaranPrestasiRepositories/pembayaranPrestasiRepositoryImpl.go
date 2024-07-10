package pembayaranprestasirepositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	pembayaranprestasimodel "github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranPrestasiModel/breakdown"
)

func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) DetailPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *[]pembayaranprestasimodel.PembayaranPrestasi, requestPembayaranPrestasi *pembayaranprestasimodel.RequestPembayaranPrestasi,totalCount *int64) error {
	pembayaranPrestasiRepositoryImpl.DB.Where("JENIS_PENGADAAN = ?", requestPembayaranPrestasi.JENIS_PENGADAAN).Find(pembayaranPrestasi).Count(totalCount)
	if err := pembayaranPrestasiRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Where("JENIS_PENGADAAN = ?", requestPembayaranPrestasi.JENIS_PENGADAAN).Find(pembayaranPrestasi).Error; err != nil {
		log.Println("pembayaranPrestasiRepositoryImpl.DB.Find(pembayaranPrestasi).Error")
		return err
	}
	return nil
}
func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) PutPembayaranPrestasi(c *fiber.Ctx, pembayaranPrestasi *pembayaranprestasimodel.PembayaranPrestasi) error {
	var whereQuery string
	if pembayaranPrestasi.NILAI_PENGADAAN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN IS NULL", pembayaranPrestasi.NAMA_PENGADAAN, pembayaranPrestasi.JENIS_PENGADAAN)
	} else {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN = '%s'", pembayaranPrestasi.NAMA_PENGADAAN, pembayaranPrestasi.JENIS_PENGADAAN, pembayaranPrestasi.NILAI_PENGADAAN)
	}
	updateProsesPengadaanModel := pembayaranPrestasiRepositoryImpl.DB.Where(whereQuery).Updates(pembayaranPrestasi)
	if updateProsesPengadaanModel.RowsAffected < 1 {
		return errors.New("data not found")
	}
	if err := updateProsesPengadaanModel.Error; err != nil {
		log.Println("updateProsesPengadaanModel.Error; err != nil")
		return err
	}
	return nil
}

func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) DetailBreakdownPembayaranPrestasi(c *fiber.Ctx, breakdownPembayaraanPrestasi *[]breakdown.BreakdownPembayaranPrestasi, breakdownRequestBreakdownPembayaranPrestasi *breakdown.RequestBreakdownPembayaranPrestasi,totalCount *int64) error {
	var whereQuery string
	if breakdownRequestBreakdownPembayaranPrestasi.NILAI_PENGADAAN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN IS NULL", breakdownRequestBreakdownPembayaranPrestasi.NAMA_PENGADAAN, breakdownRequestBreakdownPembayaranPrestasi.JENIS_PENGADAAN)
	} else {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and NILAI_PENGADAAN = '%s'", breakdownRequestBreakdownPembayaranPrestasi.NAMA_PENGADAAN, breakdownRequestBreakdownPembayaranPrestasi.JENIS_PENGADAAN, breakdownRequestBreakdownPembayaranPrestasi.NILAI_PENGADAAN)
	}
	pembayaranPrestasiRepositoryImpl.DB.Where(whereQuery).Find(breakdownPembayaraanPrestasi).Count(totalCount)
	if err := pembayaranPrestasiRepositoryImpl.DB.Where(whereQuery).Find(breakdownPembayaraanPrestasi).Error; err != nil {
		log.Println("pembayaranPrestasiRepositoryImpl.DB.Find(breakdownPembayaraanPrestasi).Error")
		return err
	}
	return nil
}
func (pembayaranPrestasiRepositoryImpl *PembayaranPrestasiRepositoryImpl) PutBreakdownPembayaranPrestasi(c *fiber.Ctx, breakdownRequestPutPembayaraanPrestasi *breakdown.RequestPutBreakdownPembayaranPrestasi) error {
	var whereQuery string
	if breakdownRequestPutPembayaraanPrestasi.NILAI_PENGADAAN == "" && breakdownRequestPutPembayaraanPrestasi.TERMIN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and TERMIN IS NULL and NILAI_PENGADAAN IS NULL", breakdownRequestPutPembayaraanPrestasi.NAMA_PENGADAAN, breakdownRequestPutPembayaraanPrestasi.JENIS_PENGADAAN)
	} else if breakdownRequestPutPembayaraanPrestasi.NILAI_PENGADAAN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and TERMIN = '%s' and NILAI_PENGADAAN IS NULL", breakdownRequestPutPembayaraanPrestasi.NAMA_PENGADAAN, breakdownRequestPutPembayaraanPrestasi.JENIS_PENGADAAN, breakdownRequestPutPembayaraanPrestasi.TERMIN)
	} else if breakdownRequestPutPembayaraanPrestasi.TERMIN == "" {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and TERMIN IS NULL and NILAI_PENGADAAN = '%s'", breakdownRequestPutPembayaraanPrestasi.NAMA_PENGADAAN, breakdownRequestPutPembayaraanPrestasi.JENIS_PENGADAAN, breakdownRequestPutPembayaraanPrestasi.NILAI_PENGADAAN)
	} else {
		whereQuery = fmt.Sprintf("NAMA_PENGADAAN = '%s' and JENIS_PENGADAAN = '%s' and TERMIN = '%s' and NILAI_PENGADAAN = '%s'", breakdownRequestPutPembayaraanPrestasi.NAMA_PENGADAAN, breakdownRequestPutPembayaraanPrestasi.JENIS_PENGADAAN, breakdownRequestPutPembayaraanPrestasi.TERMIN, breakdownRequestPutPembayaraanPrestasi.NILAI_PENGADAAN)
	}
	updateProsesPengadaanModel := pembayaranPrestasiRepositoryImpl.DB.Where(whereQuery).Updates(breakdownRequestPutPembayaraanPrestasi)
	if updateProsesPengadaanModel.RowsAffected < 1 {
		log.Println("data not found")
		return errors.New("data not found")
	}
	if err := updateProsesPengadaanModel.Error; err != nil {
		log.Println("updateProsesPengadaanModel.Error; err")
		return err
	}
	return nil
}
