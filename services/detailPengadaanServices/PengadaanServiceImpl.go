package detailpengadaanservices

import (
	"fmt"
	"log"

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
		log.Printf("error PengadaanFilterRepository.IndexPengadaan %v", err)
		return dataDetailPengadaan, err
	}

	return dataDetailPengadaan, nil
}
func (repository *PengadaanServiceImpl) IndexStatus(c *fiber.Ctx) ([]detailmodel.Status, error) {
	dataListStatus, err := repository.PengadaanFilterRepository.IndexStatus(c)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.IndexStatus %v", err)
		return dataListStatus, err
	}

	return dataListStatus, nil
}
func (repository *PengadaanServiceImpl) IndexType(c *fiber.Ctx) ([]detailmodel.Type, error) {
	dataListType, err := repository.PengadaanFilterRepository.IndexType(c)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.IndexType %v", err)
		return dataListType, err
	}

	return dataListType, nil
}
func (repository *PengadaanServiceImpl) FilterPengadaan(c *fiber.Ctx, filter map[string]string) ([]detailmodel.Pengadaan, error) {
	dataFilterDetailPengadaan, err := repository.PengadaanFilterRepository.FilterPengadaan(c, filter)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.FilterPengadaan %v", err)
		return dataFilterDetailPengadaan, err
	}

	return dataFilterDetailPengadaan, nil
}
func (repository *PengadaanServiceImpl) SumPengadaan(c *fiber.Ctx, SUM string, GROUP_BY string, WHERE_KEY string, WHERE_VALUE string) ([]detailmodel.DataResultSumPengadaan, error) {
	var sumSelectStringDetailPengadaan = fmt.Sprintf("SELECT sum(%v) AS NILAI_PENGADAAN_HASIL,%v FROM PENGADAAN WHERE %v = '%v' GROUP BY %v", SUM, GROUP_BY, WHERE_KEY, WHERE_VALUE, GROUP_BY)
	dataFilterDetailPengadaan, err := repository.PengadaanFilterRepository.SumPengadaan(c, sumSelectStringDetailPengadaan)
	if err != nil {
		log.Printf("error PengadaanFilterRepository.FilterPengadaan %v", err)
		return dataFilterDetailPengadaan, err
	}

	return dataFilterDetailPengadaan, nil
}
