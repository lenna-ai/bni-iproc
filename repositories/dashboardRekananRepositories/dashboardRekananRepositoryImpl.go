package dashboardrekananrepositories

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	pegadaanmodel "github.com/lenna-ai/bni-iproc/models/pegadaanModel"
)

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) Rekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaVendor string,rekananData *[]map[string]any,totalCount *int64) error  {
	var whereQuery string

	if strings.ToLower(param) == "non it"{
		whereQuery = "where p.jenis_pengadaan not in ('IT','Premises')"
	}else{
		whereQuery = fmt.Sprintf(`where lower(p.jenis_pengadaan) = '%v'`, strings.ToLower(param))
	}
	if filterNamaVendor != "" {
		whereQuery += fmt.Sprintf(" and NAMA_VENDOR LIKE '%%%s%%'", filterNamaVendor)
	}
	query := fmt.Sprintf(`SELECT p.NAMA_VENDOR, COUNT(*) jumlah_data_vendor FROM PENGADAAN_FILTER p %v GROUP BY p.NAMA_VENDOR`,whereQuery)
	if usePagination {
		dashboardRekananRepositoryImpl.DB.Raw(query).Count(totalCount)
	
		pageSize, offset := gormhelpers.PaginateRaw(c)
		paginatedQuery := query + " OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
		if err := dashboardRekananRepositoryImpl.DB.Raw(paginatedQuery,pageSize, offset).Scan(rekananData).Error; err != nil {
			log.Printf("dashboardRekananRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Table(PEMBAYARAN p).Select(p.NAMA_VENDOR ,COUNT(p.NAMA_PEKERJAAN) as calculate_job_name, sum(p.NILAI_KONTRAK) AS total_pekerjaan).Group(p.NAMA_VENDOR).Where(p.JENIS_PENGADAAN = ?,param).Find(rekananData).Error")
			return err
		}
	}else{
		if err := dashboardRekananRepositoryImpl.DB.Raw(query).Scan(rekananData).Error; err != nil {
			log.Printf("dashboardRekananRepositoryImpl.DB.Raw(query).Scan(rekananData).Error")
			return err
		}
	}
	return nil
}

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) BreakdownRekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaPekerjaan string,breakdownRekananData *[]pegadaanmodel.PengadaanFilter, totalCount *int64) error  {
	var whereQuery string 
	whereQuery = fmt.Sprintf("NAMA_VENDOR = '%v'",param)
	if filterNamaPekerjaan != "" {
		whereQuery += fmt.Sprintf(" and NAMA_PEKERJAAN LIKE '%%%s%%'", filterNamaPekerjaan)
	}

	if usePagination {
		dashboardRekananRepositoryImpl.DB.Model(breakdownRekananData).Where(whereQuery).Count(totalCount)
		if err := dashboardRekananRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Where(whereQuery).Find(breakdownRekananData).Error; err != nil {
			log.Printf("dashboardRekananRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Where(NAMA_VENDOR = ?,param).Find(breakdownRekananData).Error %v\n ", err)
			return err
		}
	}else{
		if err := dashboardRekananRepositoryImpl.DB.Where(whereQuery).Find(breakdownRekananData).Error; err != nil {
			log.Printf("dashboardRekananRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Where(NAMA_VENDOR = ?,param).Find(breakdownRekananData).Error %v\n ", err)
			return err
		}
	}
	return nil
}