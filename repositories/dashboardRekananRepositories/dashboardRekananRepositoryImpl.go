package dashboardrekananrepositories

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	dashboardmodel "github.com/lenna-ai/bni-iproc/models/dashboardModel"
)

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) Rekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaVendor string,rekananData *[]map[string]any,totalCount *int64) error  {
	var whereQuery string

	if strings.ToLower(param) == "non it"{
		whereQuery = "where p1.jenis_pengadaan not in ('IT','Premises')"
	}else{
		whereQuery = fmt.Sprintf(`where lower(p1.jenis_pengadaan) = '%v'`, strings.ToLower(param))
	}
	if filterNamaVendor != "" {
		whereQuery += fmt.Sprintf(" and NAMA_VENDOR LIKE '%%%s%%'", filterNamaVendor)
	}
	
	query := fmt.Sprintf(`select p1.nama_vendor, 
				(select count(*) from 
				(select count(*)
				from pengadaan p2
				where p2.nama_vendor = p1.nama_vendor
				group by p2.procurement_id))
				as jumlah_pengadaan,
				(select sum(x.nilai_spk) from 
				(select count(procurement_id) as total, p2.nilai_spk
				from pengadaan p2
				where p2.nama_vendor = p1.nama_vendor
				group by p2.procurement_id, p2.nilai_spk) x
				) as nilai_kontrak
				from pengadaan p1
				%v
				and p1.nama_vendor is not null
				group by p1.nama_vendor
	`,whereQuery)
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

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) BreakdownRekanan(c *fiber.Ctx,usePagination bool,param string,filterNamaPekerjaan string,breakdownRekananData *[]dashboardmodel.DashboardRekanan, totalCount *int64) error  {
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