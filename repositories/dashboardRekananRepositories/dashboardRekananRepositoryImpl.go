package dashboardrekananrepositories

import (
	"github.com/gofiber/fiber/v2"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
)

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) Rekanan(c *fiber.Ctx,param string,rekananData *[]map[string]any) error  {
	if err := dashboardRekananRepositoryImpl.DB.Table("PEMBAYARAN p").Select("p.NAMA_VENDOR ,COUNT(p.NAMA_PEKERJAAN) as calculate_job_name, sum(p.NILAI_KONTRAK) AS total_pekerjaan").Group("p.NAMA_VENDOR").Where("p.JENIS_PENGADAAN = ?",param).Find(rekananData).Error; err != nil {
		return err
	}
	return nil
}

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) BreakdownRekanan(c *fiber.Ctx,param string,breakdownRekananData *[]dashboardmodel.DashboardRekanan) error  {
	if err := dashboardRekananRepositoryImpl.DB.Where("NAMA_VENDOR = ?",param).Find(breakdownRekananData).Error; err != nil {
		return err
	}
	return nil
}