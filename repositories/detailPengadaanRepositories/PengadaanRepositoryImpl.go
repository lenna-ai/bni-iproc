package detailpengadaanrepositories

import (
	"log"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	detailmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
)

func (repository *PengadaanRepositoryImpl) FilterPengadaanUmum(c *fiber.Ctx, stringWhere string,totalCount *int64) ([]detailmodel.PengadaanFilter, error) {
	dataFilterDetailPengadaan := new([]detailmodel.PengadaanFilter)
	repository.DB.Where(stringWhere).Find(dataFilterDetailPengadaan).Count(totalCount)
	if err := repository.DB.Scopes(gormhelpers.Paginate(c)).Where(stringWhere).Find(dataFilterDetailPengadaan).Error; err != nil {
		log.Printf("error PengadaanRepositoryImpl.FilterPengadaan %v\n ", err)
		return *dataFilterDetailPengadaan, err
	}

	return *dataFilterDetailPengadaan, nil
}


func (repository *PengadaanRepositoryImpl) FilterPengadaanMonitoringPengadaan(c *fiber.Ctx, stringWhere string,totalCount *int64) ([]detailmodel.PengadaanFilter, error) {
	dataFilterDetailPengadaan := new([]detailmodel.PengadaanFilter)
	
	// Hitung total jumlah data tanpa pagination
	err := repository.DB.Joins("LEFT JOIN MONITORING_PROSES_PENGADAAN_NEW mppn ON PENGADAAN_FILTER.PROCUREMENT_ID = mppn.PROCUREMENT_ID").Where(stringWhere).Model(&detailmodel.PengadaanFilter{}).Count(totalCount).Error
	if err != nil {
	log.Printf("error PengadaanRepositoryImpl.FilterPengadaan Count %v\n", err)
		return *dataFilterDetailPengadaan, err
	}
	
	// Ambil data dengan pagination
	err = repository.DB.Model(dataFilterDetailPengadaan).Preload("MonitoringProses").Joins("LEFT JOIN MONITORING_PROSES_PENGADAAN_NEW mppn ON PENGADAAN_FILTER.PROCUREMENT_ID = mppn.PROCUREMENT_ID").Scopes(gormhelpers.Paginate(c)).Where(stringWhere).Find(dataFilterDetailPengadaan).Error
	if err != nil {
	log.Printf("error PengadaanRepositoryImpl.FilterPengadaan %v\n", err)
	return *dataFilterDetailPengadaan, err
	}
	
	return *dataFilterDetailPengadaan, nil
	}

func (repository *PengadaanRepositoryImpl) IndexPengadaan(c *fiber.Ctx) ([]detailmodel.Pengadaan, error) {
	dataDetailPengadaan := new([]detailmodel.Pengadaan)
	if err := repository.DB.Find(dataDetailPengadaan).Error; err != nil {
		log.Printf("error PengadaanRepositoryImpl.IndexPengadaan.Find(dataDetailPengadaan).Error %v \n", err)
		return *dataDetailPengadaan, err
	}

	return *dataDetailPengadaan, nil
}

func (repository *PengadaanRepositoryImpl) IndexStatus(c *fiber.Ctx) ([]detailmodel.Status, error) {
	dataListStatus := new([]detailmodel.Status)
	if err := repository.DB.Find(dataListStatus).Error; err != nil {
		log.Printf("error .PengadaanRepositoryImpl.IndexStatus.Find(dataListStatus).Error %v\n", err)
		return *dataListStatus, err
	}

	return *dataListStatus, nil
}

func (repository *PengadaanRepositoryImpl) IndexType(c *fiber.Ctx) ([]detailmodel.Type, error) {
	dataListType := new([]detailmodel.Type)
	if err := repository.DB.Find(dataListType).Error; err != nil {
		log.Printf("error PengadaanRepositoryImpl.IndexType.Find(dataListType).Error %v\n", err)
		return *dataListType, err
	}

	return *dataListType, nil
}

func (repository *PengadaanRepositoryImpl) SumPengadaan(c *fiber.Ctx, sumSelectStringDetailPengadaan string) ([]detailmodel.DataResultSumPengadaan, error) {
	dataSumDetailPengadaan := new([]detailmodel.DataResultSumPengadaan)
	err := repository.DB.Raw(sumSelectStringDetailPengadaan).Scan(&dataSumDetailPengadaan).Error
	if err != nil {
		log.Printf("error PengadaanRepositoryImpl.SumPengadaan.Scan(dataSumDetailPengadaan).Error %v\n", err)
		return *dataSumDetailPengadaan, err
	}
	return *dataSumDetailPengadaan, nil
}
