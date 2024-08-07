package detailpengadaanservices

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	detailmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
)

func (repository *PengadaanServiceImpl) IndexPengadaan(c *fiber.Ctx) ([]detailmodel.Pengadaan, error) {
	dataDetailPengadaan, err := repository.PengadaanFilterRepository.IndexPengadaan(c)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.IndexPengadaan %v\n", err)
		return dataDetailPengadaan, err
	}

	return dataDetailPengadaan, nil
}
func (repository *PengadaanServiceImpl) IndexStatus(c *fiber.Ctx) ([]detailmodel.Status, error) {
	dataListStatus, err := repository.PengadaanFilterRepository.IndexStatus(c)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.IndexStatus %v\n", err)
		return dataListStatus, err
	}

	return dataListStatus, nil
}
func (repository *PengadaanServiceImpl) IndexType(c *fiber.Ctx) ([]detailmodel.Type, error) {
	dataListType, err := repository.PengadaanFilterRepository.IndexType(c)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.IndexType %v\n", err)
		return dataListType, err
	}

	return dataListType, nil
}
func (repository *PengadaanServiceImpl) FilterPengadaan(c *fiber.Ctx,usePagination bool, filter map[string]string,totalCount *int64) ([]detailmodel.PengadaanFilter, error) {
	var stringWhere string
	var loopFilter int
	for k, v := range filter {
		if loopFilter < len(filter)-1 {
			stringWhere += fmt.Sprintf("%v = '%v' AND ", k, v)
			loopFilter++
		} else {
			stringWhere += fmt.Sprintf("%v = '%v'", k, v)
		}
	}
	var dataFilterDetailPengadaan []detailmodel.PengadaanFilter

	if c.Query("filter_for") == "monitoring_proses_pengadaan"{
		filterPengadaanMonitoringPengadaan, err := repository.PengadaanFilterRepository.FilterPengadaanMonitoringPengadaan(c,usePagination, stringWhere, totalCount)
		dataFilterDetailPengadaan = filterPengadaanMonitoringPengadaan
		if err != nil {
			log.Printf("error PengadaanFilterRepository.FilterPengadaanMonitoringPengadaan %v\n", err)
			return dataFilterDetailPengadaan, err
		}
	}else{
		filterPengadaanUmum, err := repository.PengadaanFilterRepository.FilterPengadaanUmum(c,usePagination, stringWhere, totalCount)
		dataFilterDetailPengadaan = filterPengadaanUmum
		if err != nil {
			log.Printf("error PengadaanFilterRepository.FilterPengadaanUmum %v\n", err)
			return dataFilterDetailPengadaan, err
		}
	}

	return dataFilterDetailPengadaan, nil
}
func (repository *PengadaanServiceImpl) SumPengadaan(c *fiber.Ctx, SUM1 string, SUM2 string, GROUP_BY string, WHERE_KEY string, WHERE_VALUE string, WHERE_SYMBOL string) ([]detailmodel.DataResultSumPengadaan, error) {
	whereKeySplit := strings.Split(WHERE_KEY, "-")
	whereValueSplit := strings.Split(WHERE_VALUE, "-")
	whereSymbolSplit := strings.Split(WHERE_SYMBOL, "-")

	var tempWhereClauses = splitStringWhere(whereKeySplit, whereValueSplit, whereSymbolSplit)

	dataFilterDetailPengadaan, err := repository.PengadaanFilterRepository.SumPengadaan(c, SUM1, SUM2, GROUP_BY, tempWhereClauses)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.FilterPengadaan %v\n", err)
		return dataFilterDetailPengadaan, err
	}

	return dataFilterDetailPengadaan, nil
}

func (repository *PengadaanServiceImpl) EfisiensiPengadaan(c *fiber.Ctx, estimasi_nilai_pengadaan int, nilai_spk int) (resultSisaAnggaran int, resultEfisiensi float64) {
	resultSisaAnggaran = estimasi_nilai_pengadaan - nilai_spk
	resultEfisiensi = (float64(resultSisaAnggaran) / float64(estimasi_nilai_pengadaan)) * 100

	return resultSisaAnggaran, resultEfisiensi
}

func (repository *PengadaanServiceImpl) DynamicPengadaan(c *fiber.Ctx,pagination bool,table string, filter map[string]string, dataResult *[]map[string]any,totalCount *int64) error {
	var stringWhere string
	var loopFilter int
	for k, v := range filter {
		if loopFilter < len(filter)-1 {
			stringWhere += fmt.Sprintf("%v = '%v' AND ", k, v)
			loopFilter++
		} else {
			stringWhere += fmt.Sprintf("%v = '%v'", k, v)
		}
	}

	if err := repository.PengadaanFilterRepository.DynamicPengadaan(c,pagination,table,filter,stringWhere,dataResult,totalCount); err != nil {
		return err
	}
	return nil
}
