package dashboardpembayarancontroller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/helpers"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	dashboardpembayaraanservices "github.com/lenna-ai/bni-iproc/services/dashboardPembayaraanServices"
)

func NewPembayaranMonitoringController(pembayaranMonitoringServices dashboardpembayaraanservices.PembayaranMonitoringServices) *PembayaranMonitoringControllerImpl {
	return &PembayaranMonitoringControllerImpl{
		PembayaranMonitoringServices: pembayaranMonitoringServices,
	}
}

func (pembayaranMonitoringControllerImpl *PembayaranMonitoringControllerImpl) IndexRekananPembayaran(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	jenisPengadaan := c.Query("JENIS_PENGADAAN")
	var totalCount = new(int64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	pembayaran, err := pembayaranMonitoringControllerImpl.PembayaranMonitoringServices.IndexRekananPembayaranService(c, jenisPengadaan,totalCount)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexStatus %v\n ", err)
		return helpers.ResultFailedJsonApi(c, pembayaran, err.Error())
	}

	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,pembayaran))
}

func (pembayaranMonitoringControllerImpl *PembayaranMonitoringControllerImpl) FilterPengadaan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	jenisPengadaan := c.Query("filter")
	var totalCount = new(int64)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	pembayaran, err := pembayaranMonitoringControllerImpl.PembayaranMonitoringServices.FilterPengadaan(c, jenisPengadaan,totalCount)
	if err != nil {
		log.Printf("error PengadaanFilterService.IndexStatus %v\n ", err)
		return helpers.ResultFailedJsonApi(c, pembayaran, err.Error())
	}

	
	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,pembayaran))
}

func (pembayaranMonitoringControllerImpl *PembayaranMonitoringControllerImpl) IndexPembayaran(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	// jenisPengadaan := c.Query("filter")

	// pembayaran, err := pembayaranMonitoringControllerImpl.PembayaranMonitoringServices.FilterPengadaan(c, jenisPengadaan)
	// if err != nil {
	// 	log.Printf("error PengadaanFilterService.IndexStatus %v\n ", err)
	// 	return helpers.ResultFailedJsonApi(c, pembayaran, err.Error())
	// }

	// return helpers.ResultSuccessJsonApi(c, pembayaran)
	return nil
}
