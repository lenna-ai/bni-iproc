package detailpengadaanservices

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	detailmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
)

func NewDetailPengadaanService(repository detailpengadaanrepositories.PengadaanRepository) *PengadaanServiceImpl {
	return &PengadaanServiceImpl{
		PengadaanFilterRepository: repository,
	}
}

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
func (repository *PengadaanServiceImpl) FilterPengadaan(c *fiber.Ctx, filter map[string]string) ([]detailmodel.Pengadaan, error) {
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
	dataFilterDetailPengadaan, err := repository.PengadaanFilterRepository.FilterPengadaan(c, stringWhere)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.FilterPengadaan %v\n", err)
		return dataFilterDetailPengadaan, err
	}

	return dataFilterDetailPengadaan, nil
}
func (repository *PengadaanServiceImpl) SumPengadaan(c *fiber.Ctx, SUM1 string, SUM2 string, GROUP_BY string, WHERE_KEY string, WHERE_VALUE string, WHERE_SYMBOL string) ([]detailmodel.DataResultSumPengadaan, error) {
	whereKeySplit := strings.Split(WHERE_KEY, "-")
	whereValueSplit := strings.Split(WHERE_VALUE, "-")
	whereSymbolSplit := strings.Split(WHERE_SYMBOL, "-")

	var tempWhereClauses = splitStringWhere(whereKeySplit, whereValueSplit, whereSymbolSplit)

	var sumSelectStringDetailPengadaan = fmt.Sprintf("SELECT sum(%v) AS ESTIMASI_NILAI_PENGADAAN, SUM(%v) AS NILAI_SPK,%v FROM PENGADAAN WHERE %v GROUP BY %v", SUM1, SUM2, GROUP_BY, tempWhereClauses, GROUP_BY)
	dataFilterDetailPengadaan, err := repository.PengadaanFilterRepository.SumPengadaan(c, sumSelectStringDetailPengadaan)
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
