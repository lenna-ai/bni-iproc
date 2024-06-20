package dashboardpembayaranservices

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	pembayaranmodel "github.com/lenna-ai/bni-iproc/models/pembayaranModel"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
)

func NewPembayaranMonitoringServices(PembayaranMonitoringRepository dashboardpembayaranrepositories.PembayaranMonitoringRepository) *PembayaranMonitoringServicesImpl {
	return &PembayaranMonitoringServicesImpl{
		PembayaranMonitoringRepository: PembayaranMonitoringRepository,
	}
}

func (pembayaranMonitoringServicesImpl *PembayaranMonitoringServicesImpl) IndexRekananPembayaranService(c *fiber.Ctx, jenisPengadaan string) ([]formatters.IndexPembayaranMonitor, error) {
	pembayaran, err := pembayaranMonitoringServicesImpl.PembayaranMonitoringRepository.IndexRekananPembayaranMonitor(c, jenisPengadaan)
	if err != nil {
		log.Printf("error pembayaranMonitoringServicesImpl.PembayaranMonitoringRepository.IndexPembayaranMonitor %v\n", err)
		return pembayaran, err
	}
	return pembayaran, nil
}

func (pembayaranMonitoringServicesImpl *PembayaranMonitoringServicesImpl) FilterPengadaan(c *fiber.Ctx, jenisPengadaan string) ([]pembayaranmodel.Pembayaran, error) {
	splitJenisPengadaan := strings.Split(jenisPengadaan, ",")
	var queryStringWhere string
	var countSplitVSJP int
	for _, vsjp := range splitJenisPengadaan {
		splitVSJP := strings.Split(vsjp, "=")
		for k, vSplitVSJP := range splitVSJP {
			if k%2 == 0 {
				queryStringWhere += vSplitVSJP + " = "
			} else {
				if countSplitVSJP+2 <= len(splitVSJP) {
					queryStringWhere += fmt.Sprintf("'%v' AND ", vSplitVSJP)
				} else {
					queryStringWhere += fmt.Sprintf("'%v'", vSplitVSJP)
				}
				countSplitVSJP++
			}
		}
	}

	pembayaran, err := pembayaranMonitoringServicesImpl.PembayaranMonitoringRepository.FilterPengadaan(c, queryStringWhere)
	if err != nil {
		log.Printf("error pembayaranMonitoringServicesImpl.PembayaranMonitoringRepository.IndexPembayaranMonitor %v\n", err)
		return pembayaran, err
	}
	return pembayaran, nil
}
