package monitoringrepositories

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
)

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) DetailPembayaranRutin(c *fiber.Ctx, pembayaranRutin *[]pembayaranrutinmodel.PembayaranRutin) error {
	if err := pembayaranRutinRepositoryImpl.DB.Find(pembayaranRutin).Error; err != nil {
		return err
	}
	return nil
}

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) PutPembayaranRutin(c *fiber.Ctx, pembayaranRutin *pembayaranrutinmodel.PembayaranRutin) error {
	var whereQuery string
	if pembayaranRutin.NilaiPengadaanHasil == "" {
		whereQuery = fmt.Sprintf("NAMA = '%s' and NILAI_PENGADAAN_HASIL IS NULL", pembayaranRutin.Nama)
	} else {
		whereQuery = fmt.Sprintf("NAMA = '%s' and NILAI_PENGADAAN_HASIL = '%s'", pembayaranRutin.Nama, pembayaranRutin.NilaiPengadaanHasil)
	}
	updatePembayaranRutinRepository := pembayaranRutinRepositoryImpl.DB.Where(whereQuery).Updates(pembayaranRutin)
	if updatePembayaranRutinRepository.RowsAffected < 1 {
		return errors.New("Data not found")
	}
	if err := updatePembayaranRutinRepository.Error; err != nil {
		return err
	}
	return nil
}

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) DetailBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *[]pembayaranrutinmodel.BreakdownPembayaranRutin) error {
	if err := pembayaranRutinRepositoryImpl.DB.Find(breakdownPembayaranRutin).Error; err != nil {
		return err
	}
	return nil
}

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) PutBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *pembayaranrutinmodel.BreakdownPembayaranRutin) error {
	updatePembayaranRutinRepository := pembayaranRutinRepositoryImpl.DB.Where("MONITORING_PEMBAYARAN_RUTIN_ID = ?", breakdownPembayaranRutin.MonitoringPembayaranRutinId).Updates(breakdownPembayaranRutin)
	if updatePembayaranRutinRepository.RowsAffected < 1 {
		return errors.New("Data not found")
	}
	if err := updatePembayaranRutinRepository.Error; err != nil {
		return err
	}
	return nil
}
