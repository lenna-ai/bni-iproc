package dashboardrekananrepositories

import (
	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
)

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) Rekanan(c *fiber.Ctx,param string,rekananData *[]map[string]any,totalCount *int64) error  {
	dashboardRekananRepositoryImpl.DB.Table("PEMBAYARAN p").Select("p.NAMA_VENDOR ,COUNT(p.NAMA_PEKERJAAN) as calculate_job_name, sum(p.NILAI_KONTRAK) AS total_pekerjaan").Group("p.NAMA_VENDOR").Where("p.JENIS_PENGADAAN = ?",param).Count(totalCount)
	if err := dashboardRekananRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Table("PEMBAYARAN p").Select("p.NAMA_VENDOR ,COUNT(p.NAMA_PEKERJAAN) as calculate_job_name, sum(p.NILAI_KONTRAK) AS total_pekerjaan").Group("p.NAMA_VENDOR").Where("p.JENIS_PENGADAAN = ?",param).Find(rekananData).Error; err != nil {
		return err
	}
	return nil
}

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) BreakdownRekanan(c *fiber.Ctx,param string,breakdownRekananData *[]dashboardmodel.DashboardRekanan, totalCount *int64) error  {
	dashboardRekananRepositoryImpl.DB.Model(breakdownRekananData).Where("NAMA_VENDOR = ?", param).Count(totalCount)
	if err := dashboardRekananRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Where("NAMA_VENDOR = ?",param).Find(breakdownRekananData).Error; err != nil {
		return err
	}
	return nil
}