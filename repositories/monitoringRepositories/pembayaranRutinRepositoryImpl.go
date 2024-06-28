package monitoringrepositories

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
)

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) PutPembayaranRutin(c *fiber.Ctx, pembayaranRutin *pembayaranrutinmodel.PembayaranRutin) error {
	updatePembayaranRutinRepository := pembayaranRutinRepositoryImpl.DB.Where("NAMA = ? and NILAI_PENGADAAN_HASIL = ?", pembayaranRutin.Nama, pembayaranRutin.NilaiPengadaanHasil).Updates(pembayaranRutin)
	if updatePembayaranRutinRepository.RowsAffected < 1 {
		return errors.New("Data not found")
	}
	if err := updatePembayaranRutinRepository.Error; err != nil {
		return err
	}
	return nil
}
