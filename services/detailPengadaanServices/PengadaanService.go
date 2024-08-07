package detailpengadaanservices

import (
	"github.com/gofiber/fiber/v2"
	detailmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
)

type PengadaanServiceImpl struct {
	PengadaanFilterRepository detailpengadaanrepositories.PengadaanRepository
}

type PengadaanService interface {
	IndexPengadaan(c *fiber.Ctx) ([]detailmodel.Pengadaan, error)
	IndexStatus(c *fiber.Ctx) ([]detailmodel.Status, error)
	IndexType(c *fiber.Ctx) ([]detailmodel.Type, error)
	FilterPengadaan(c *fiber.Ctx,usePagination bool, filter map[string]string,totalCount *int64) ([]detailmodel.PengadaanFilter, error)
	SumPengadaan(c *fiber.Ctx, SUM1 string, SUM2 string, GROUP_BY string, WHERE_KEY string, WHERE_VALUE string, WHERE_AND string) ([]detailmodel.DataResultSumPengadaan, error)
	EfisiensiPengadaan(c *fiber.Ctx, estimasi_nilai_pengadaan int, nilai_spk int) (resultSisaAnggaran int, resultEfisiensi float64)
	DynamicPengadaan(c *fiber.Ctx,pagination bool,table string, filter map[string]string,dataResult *[]map[string]any,totalCount *int64) error
}

func NewDetailPengadaanService(repository detailpengadaanrepositories.PengadaanRepository) *PengadaanServiceImpl {
	return &PengadaanServiceImpl{
		PengadaanFilterRepository: repository,
	}
}