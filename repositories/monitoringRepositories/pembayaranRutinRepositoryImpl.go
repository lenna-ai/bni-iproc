package monitoringrepositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	pembayaranrutinmodel "github.com/lenna-ai/bni-iproc/models/pembayaranRutinModel"
)

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) DetailPembayaranRutin(c *fiber.Ctx, pembayaranRutin *[]pembayaranrutinmodel.PembayaranRutin,totalCount *int64) error {
	pembayaranRutinRepositoryImpl.DB.Find(pembayaranRutin).Count(totalCount)
	if err := pembayaranRutinRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Find(pembayaranRutin).Error; err != nil {
		log.Println("pembayaranRutinRepositoryImpl.DB.Find(pembayaranRutin).Error; err")
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
		log.Println("updatePembayaranRutinRepository.RowsAffected < 1")
		return errors.New("data not found")
	}
	if err := updatePembayaranRutinRepository.Error; err != nil {
		log.Println("updatePembayaranRutinRepository.Error; err != nil")
		return err
	}
	return nil
}

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) DetailBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *[]pembayaranrutinmodel.BreakdownPembayaranRutin,totalCount *int64) error {
	pembayaranRutinRepositoryImpl.DB.Find(breakdownPembayaranRutin).Count(totalCount)
	if err := pembayaranRutinRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Find(breakdownPembayaranRutin).Error; err != nil {
		log.Println("epositoryImpl.DB.Find(breakdownPembayaranRutin).Error; err")
		return err
	}
	return nil
}

func (pembayaranRutinRepositoryImpl *PembayaranRutinRepositoryImpl) PutBreakdownPembayaranRutin(c *fiber.Ctx, breakdownPembayaranRutin *pembayaranrutinmodel.BreakdownPembayaranRutin) error {
	updatePembayaranRutinRepository := pembayaranRutinRepositoryImpl.DB.Where("MONITORING_PEMBAYARAN_RUTIN_ID = ?", breakdownPembayaranRutin.MonitoringPembayaranRutinId).Updates(breakdownPembayaranRutin)
	if updatePembayaranRutinRepository.RowsAffected < 1 {
		log.Println("updatePembayaranRutinRepository.RowsAffected < 1")
		return errors.New("data not found")
	}
	if err := updatePembayaranRutinRepository.Error; err != nil {
		log.Println("datePembayaranRutinRepository.Error; err != nil")
		return err
	}
	return nil
}
