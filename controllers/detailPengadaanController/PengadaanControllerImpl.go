package detailpengadaancontroller

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
)

func NewDetailPengadaanController(DetailPengadaanFilterService detailpengadaanservices.PengadaanService) *PengadaanControllerImpl {
	return &PengadaanControllerImpl{
		PengadaanFilterService: DetailPengadaanFilterService,
	}
}

func (FilterController *PengadaanControllerImpl) IndexPengadaan(c *fiber.Ctx) error {
	dataDetailPengadaan, err := FilterController.PengadaanFilterService.IndexPengadaan(c)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexPengadaan %v", err)
		return helpers.ResultFailedJsonApi(c, dataDetailPengadaan, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataDetailPengadaan)
}

func (FilterController *PengadaanControllerImpl) FilterPengadaan(c *fiber.Ctx) error {
	filter := make(map[string]string)
	status_pengadaan := c.Query("filter")
	for _, valueSplitStatusPengadaan := range strings.Split(status_pengadaan, ",") {
		for i := 0; i < len(strings.Split(valueSplitStatusPengadaan, "="))/2; i++ {
			filter[strings.Split(valueSplitStatusPengadaan, "=")[i]] = strings.Split(valueSplitStatusPengadaan, "=")[i+1]
		}
	}
	dataFilterDetailPengadaan, err := FilterController.PengadaanFilterService.FilterPengadaan(c, filter)
	if err != nil {
		log.Printf("error PengadaanFilterService.FilterPengadaan %v", err)
		return helpers.ResultFailedJsonApi(c, dataFilterDetailPengadaan, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataFilterDetailPengadaan)
}

func (FilterController *PengadaanControllerImpl) IndexStatus(c *fiber.Ctx) error {
	dataListStatus, err := FilterController.PengadaanFilterService.IndexStatus(c)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexStatus %v", err)
		return helpers.ResultFailedJsonApi(c, dataListStatus, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataListStatus)
}

func (FilterController *PengadaanControllerImpl) IndexType(c *fiber.Ctx) error {
	dataListType, err := FilterController.PengadaanFilterService.IndexType(c)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexType %v", err)
		return helpers.ResultFailedJsonApi(c, dataListType, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataListType)
}
