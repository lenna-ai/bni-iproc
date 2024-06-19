package dashboardpembayaranservices

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/models/pembayaranModel/formatters"
	dashboardpembayaranrepositories "github.com/lenna-ai/bni-iproc/repositories/dashboardPembayaranRepositories"
)

func NewPembayaranMonitoringServices(PembayaranMonitoringRepository dashboardpembayaranrepositories.PembayaranMonitoringRepository) *PembayaranMonitoringServicesImpl {
	return &PembayaranMonitoringServicesImpl{
		PembayaranMonitoringRepository: PembayaranMonitoringRepository,
	}
}

func (pembayaranMonitoringServicesImpl *PembayaranMonitoringServicesImpl) IndexPengadaanService(c *fiber.Ctx, jenisPengadaan string) ([]formatters.IndexPembayaranMonitor, error) {
	pembayaran, err := pembayaranMonitoringServicesImpl.PembayaranMonitoringRepository.IndexPembayaranMonitor(c, jenisPengadaan)
	if err != nil {
		log.Printf("error pembayaranMonitoringServicesImpl.PembayaranMonitoringRepository.IndexPembayaranMonitor %v\n", err)
		return pembayaran, err
	}
	return pembayaran, nil
}
