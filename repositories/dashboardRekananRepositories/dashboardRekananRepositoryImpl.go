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
		// whereQuery = "where p.jenis_pengadaan not in ('IT','Premises') and p.NAMA_VENDOR is not null"
		whereQuery = "p.jenis_pengadaan not in ('IT','Premises') and p.NAMA_VENDOR is not null"
	}else{
		// whereQuery = fmt.Sprintf(`where lower(p.jenis_pengadaan) = '%v' and p.NAMA_VENDOR is not null`, strings.ToLower(param))
		whereQuery = fmt.Sprintf(`lower(p.jenis_pengadaan) = '%v' and p.NAMA_VENDOR is not null`, strings.ToLower(param))
	}
	if filterNamaVendor != "" {
		whereQuery += fmt.Sprintf(" and NAMA_VENDOR LIKE '%%%s%%'", filterNamaVendor)
	}

	if usePagination {
		// dashboardRekananRepositoryImpl.DB.Raw(query).Count(totalCount)
	
		// pageSize, offset := gormhelpers.PaginateRaw(c)
		// paginatedQuery := query + " OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
		// if err := dashboardRekananRepositoryImpl.DB.Raw(paginatedQuery,pageSize, offset).Scan(rekananData).Error; err != nil {
		// 	log.Printf("dashboardRekananRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Table(PEMBAYARAN p).Select(p.NAMA_VENDOR ,COUNT(p.NAMA_PEKERJAAN) as calculate_job_name, sum(p.NILAI_KONTRAK) AS total_pekerjaan).Group(p.NAMA_VENDOR).Where(p.JENIS_PENGADAAN = ?,param).Find(rekananData).Error")
		// 	return err
		// }
		dashboardRekananRepositoryImpl.DB.Table("PENGADAAN_FILTER p").Select("NAMA_VENDOR, NVL(SUM(p.NILAI_SPK), 0) AS nilai_kontrak,COUNT(*) AS jumlah_pengadaan_vendor").Where(whereQuery).Group("p.NAMA_VENDOR").Count(totalCount)
		if err := dashboardRekananRepositoryImpl.DB.Table("PENGADAAN_FILTER p ").Scopes(gormhelpers.Paginate(c)).Select("NAMA_VENDOR, NVL(SUM(p.NILAI_SPK), 0) AS nilai_kontrak,COUNT(*) AS jumlah_pengadaan_vendor").Where(whereQuery).Group("p.NAMA_VENDOR").Find(rekananData).Error; err!= nil{
			return err
		}
	}else{
		// if err := dashboardRekananRepositoryImpl.DB.Raw(query).Scan(rekananData).Error; err != nil {
		// 	log.Printf("dashboardRekananRepositoryImpl.DB.Raw(query).Scan(rekananData).Error")
		// 	return err
		// }
		dashboardRekananRepositoryImpl.DB.Table("PENGADAAN_FILTER p").Scopes(gormhelpers.Paginate(c)).Select("NAMA_VENDOR, NVL(SUM(p.NILAI_SPK), 0) AS nilai_kontrak,COUNT(*) AS jumlah_pengadaan_vendor").Where(whereQuery).Group("p.NAMA_VENDOR").Count(totalCount)
		if err := dashboardRekananRepositoryImpl.DB.Table("PENGADAAN_FILTER p").Select("NAMA_VENDOR, NVL(SUM(p.NILAI_SPK), 0) AS nilai_kontrak,COUNT(*) AS jumlah_pengadaan_vendor").Where(whereQuery).Group("p.NAMA_VENDOR").Find(rekananData).Error; err!= nil{
			return err
		}
	}
	return nil
}

func (dashboardRekananRepositoryImpl *DashboardRekananRepositoryImpl) BreakdownRekanan(c *fiber.Ctx,usePagination bool,param string,jenis_pengadaan string,filterNamaPekerjaan string,breakdownRekananData *[]pegadaanmodel.PengadaanFilter, totalCount *int64) error  {
	var whereQuery string 
	// whereQuery = fmt.Sprintf("STATUS IN ('Done','On Progress','revision','waiting approval')")

	if strings.ToLower(param) == "non it"{
		whereQuery = "jenis_pengadaan not in ('IT','Premises') "
	}else{
		whereQuery = fmt.Sprintf(`lower(jenis_pengadaan) = '%v' `, strings.ToLower(jenis_pengadaan))
	}
	if filterNamaPekerjaan != "" {
		whereQuery += fmt.Sprintf("and NAMA_PEKERJAAN LIKE '%%%s%%'", filterNamaPekerjaan)
	}
	whereQuery += fmt.Sprintf("AND NAMA_VENDOR = '%v' STATUS_PENGADAAN = Done",param)

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