package detailpengadaancontroller

import (
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	"github.com/lenna-ai/bni-iproc/models/pegadaanModel/formatters"
)

func (FilterController *PengadaanControllerImpl) IndexPengadaan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	dataDetailPengadaan, err := FilterController.PengadaanFilterService.IndexPengadaan(c)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexPengadaan %v \n ", err)
		return helpers.ResultFailedJsonApi(c, dataDetailPengadaan, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataDetailPengadaan)
}

func (FilterController *PengadaanControllerImpl) FilterPengadaan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	filter := make(map[string]string)
	status_pengadaan := c.Query("filter")
	for _, valueSplitStatusPengadaan := range strings.Split(status_pengadaan, ",") {
		for i := 0; i < len(strings.Split(valueSplitStatusPengadaan, "="))/2; i++ {
			filter[strings.Split(valueSplitStatusPengadaan, "=")[i]] = strings.Split(valueSplitStatusPengadaan, "=")[i+1]
		}
	}
	dataFilterDetailPengadaan, err := FilterController.PengadaanFilterService.FilterPengadaan(c, filter)
	if err != nil {
		log.Printf("error PengadaanFilterService.FilterPengadaan %v\n ", err)
		return helpers.ResultFailedJsonApi(c, dataFilterDetailPengadaan, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataFilterDetailPengadaan)
}

func (FilterController *PengadaanControllerImpl) IndexStatus(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	dataListStatus, err := FilterController.PengadaanFilterService.IndexStatus(c)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexStatus %v\n ", err)
		return helpers.ResultFailedJsonApi(c, dataListStatus, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataListStatus)
}

func (FilterController *PengadaanControllerImpl) IndexType(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	dataListType, err := FilterController.PengadaanFilterService.IndexType(c)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexType %v\n ", err)
		return helpers.ResultFailedJsonApi(c, dataListType, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataListType)
}

func (FilterController *PengadaanControllerImpl) SumPengadaan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	// var SUM_NILAI_PENGADAAN_HASIL = "NILAI_PENGADAAN_HASIL"
	// var GROUP_PENGADAAN = "JENIS_PENGADAAN"
	// var WHERE_KEY = "MATA_ANGGARAN-JENIS_PENGADAAN"
	// var WHERE_VALUE = "'Opex'-('IT','Non IT','Premises')"
	// var WHERE_SYMBOL = "=-IN"
	sum1 := c.Query("SUM1")
	sum2 := c.Query("SUM2")
	group_by := c.Query("GROUP_BY")
	where_key := c.Query("WHERE_KEY")
	where_value := c.Query("WHERE_VALUE")
	where_symbol := c.Query("WHERE_SYMBOL")

	dataFilterDetailPengadaan, err := FilterController.PengadaanFilterService.SumPengadaan(c, sum1, sum2, group_by, where_key, where_value, where_symbol)
	if err != nil {
		log.Printf("error PengadaanFilterService.FilterPengadaan %v\n ", err)
		return helpers.ResultFailedJsonApi(c, dataFilterDetailPengadaan, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, dataFilterDetailPengadaan)
}

func (FilterController *PengadaanControllerImpl) EfisiensiPengadaan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	estimasi_nilai_pengadaan, err := strconv.Atoi(c.Query("ESTIMASI_NILAI_PENGADAAN"))
	if err != nil {
		log.Printf("error ESTIMASI_NILAI_PENGADAAN converter %v\n ", err)
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	nilai_spk, err := strconv.Atoi(c.Query("NILAI_SPK"))
	if err != nil {
		log.Printf("error NILAI_SPK converter %v\n ", err)
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	resultSisaAnggaran, resultEfisiensi := FilterController.PengadaanFilterService.EfisiensiPengadaan(c, estimasi_nilai_pengadaan, nilai_spk)

	var data = formatters.EfisiensiPengadaan{
		ResultSisaAnggaran: resultSisaAnggaran,
		ResultEfisiensi:    resultEfisiensi,
	}
	return helpers.ResultSuccessJsonApi(c, data)
}
