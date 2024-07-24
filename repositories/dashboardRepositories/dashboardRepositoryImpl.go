package dashboardrepositories

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
)

func (dashboardRepositoryImpl *DashboardRepositoryImpl) TotalPengadaan(c *fiber.Ctx,dashboardModel *map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select("SUM(p.nilai_spk) as nilai_spk").Find(dashboardModel).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(SUM(p.nilai_spk) as nilai_spk).Find(dashboardModel).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) TotalPembayaran(c *fiber.Ctx,dashboardModel *map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("BREAKDOWN_MONITORING_PEMBAYARAN_RUTIN bmpr").Select("sum(BMPR.NILAI_TAGIHAN) as nilai_tagihan").Where("BMPR.STATUS_PEMBAYARAN = ?","sudah dibayarkan").Find(dashboardModel).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(BREAKDOWN_MONITORING_PEMBAYARAN_RUTIN bmpr).Select(sum(BMPR.NILAI_TAGIHAN) as nilai_tagihan).Where(BMPR.STATUS_PEMBAYARAN = ?","sudah dibayarkan).Find(dashboardModel).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) TotalVendor(c *fiber.Ctx,dashboardModel *map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("DATA_VENDOR_RESULT dvr").Select("sum(p.NILAI_SPK) as nilai_spk").Joins("right join PENGADAAN p ON p.VENDOR_ID = dvr.ID").Find(dashboardModel).Error; err != nil {
		log.Println("err := dashboardRepositoryImpl.DB.Table(DATA_VENDOR_RESULT dvr).Select(sum(p.NILAI_SPK) as nilai_spk).Joins(right join PENGADAAN p ON p.VENDOR_ID = dvr.ID).Find(dashboardModel).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error {
	query := `
	SELECT df.Kewenangan,
		   COUNT(CASE WHEN df.jenis_pengadaan = 'IT' THEN 1 END) AS total_it,
		   COUNT(CASE WHEN df.jenis_pengadaan not in ('IT','Premises') THEN 1 END) AS total_nonit,
		   COUNT(CASE WHEN df.jenis_pengadaan = 'PREMISES' THEN 1 END) AS total_premises
	FROM (
		SELECT p.*,
			   CASE
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('pemimpin departemen divisi pfa (unit pelaksana)'), LOWER('pemimpin departemen (unit pengguna)'), LOWER('wakil general manager')) THEN 'TPD-1'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('pemimpin divisi pfa (unit pelaksana)'), LOWER('pemimpin divisi/satuan (unit pengguna)'), LOWER('general manager')) THEN 'TPD-2'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('direktur yang membawahkan fungsi pengadaan'), LOWER('direktur yang membawakan fungsi manajemen risiko'), LOWER('dir. Bidang/SEVP (unit pengguna)')) THEN 'TPP-1'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('dirut'), LOWER('wadirut'), LOWER('direktur yang membawahkan fungsi pengadaan'), LOWER('direktur yang membawakan fungsi manajemen risiko'), LOWER('dir. Bidang/sevp (unit pengguna)')) THEN 'TPP-2'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) = LOWER('Rapat Direksi') THEN 'TPP-3'
					ELSE 'Not Found'
			   END AS Kewenangan
		FROM PENGADAAN p
		WHERE p.STATUS_PENGADAAN = 'Done'
	) df
	GROUP BY df.Kewenangan
	ORDER BY df.Kewenangan ASC`	
	
	if err := dashboardRepositoryImpl.DB.Raw(query).Scan(dashboardModel).Error; err != nil {
		log.Println("err := dashboardRepositoryImpl.DB.Table((?) subquery, subQuery).Select(Kewenangan, COUNT(*) as Count).Group(Kewenangan).Scan(dashboardModel).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error {

	// query := `
	// 		SELECT df.STATUS_PENGADAAN, 
	// 			COUNT(CASE WHEN df.jenis_pengadaan = 'IT' and status_pengadaan = 'On Progress' THEN 1 END) AS total_it,
	// 			COUNT(CASE WHEN df.jenis_pengadaan not in ('IT','Premises') and status_pengadaan = 'On Progress' THEN 1 END) AS total_nonit,
	// 			COUNT(CASE WHEN df.jenis_pengadaan = 'PREMISES' and status_pengadaan = 'On Progress' THEN 1 END) AS total_premises
	// 		FROM (SELECT p.* FROM PENGADAAN p) df
	// 		GROUP BY df.STATUS_PENGADAAN
	// 		ORDER BY df.STATUS_PENGADAAN ASC`

	if err := dashboardRepositoryImpl.DB.Table("STATUSPENGADAANONGOING").Find(statusPengadaan).Error; err != nil{
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.STATUS_PENGADAAN, COUNT(*) as count_pengadaan).Group(p.STATUS_PENGADAAN).Find(statusPengadaan).Error; err != nil")
		return err
	}

	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {

	if err := dashboardRepositoryImpl.DB.Table("metodeOnGoing").Find(metodePengadaan).Error; err != nil{
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.METODE ,count(*) as count_metode).Group(p.METODE).Where(p.STATUS_PENGADAAN = ?,On Progress).Find(metodePengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingKeputusan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("pendingKeputusanPengadaan").Find(metodePengadaan).Error; err != nil{
		log.Println("(dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingKeputusan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error")
		return err
	}
	
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error {
	query := `
	SELECT df.Kewenangan,
		   COUNT(CASE WHEN df.jenis_pengadaan = 'IT' THEN 1 END) AS total_it,
		   COUNT(CASE WHEN df.jenis_pengadaan not in ('IT','Premises') THEN 1 END) AS total_nonit,
		   COUNT(CASE WHEN df.jenis_pengadaan = 'PREMISES' THEN 1 END) AS total_premises
	FROM (
		SELECT p.*,
			   CASE
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('pemimpin departemen divisi pfa (unit pelaksana)'), LOWER('pemimpin departemen (unit pengguna)'), LOWER('wakil general manager')) THEN 'TPD-1'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('pemimpin divisi pfa (unit pelaksana)'), LOWER('pemimpin divisi/satuan (unit pengguna)'),LOWER('general manager')) THEN 'TPD-2'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('direktur yang membawahkan fungsi pengadaan'), LOWER('direktur yang membawakan fungsi manajemen risiko'), LOWER('dir. Bidang/SEVP (unit pengguna)')) THEN 'TPP-1'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) IN (LOWER('dirut'), LOWER('wadirut'), LOWER('direktur yang membawahkan fungsi pengadaan'), LOWER('direktur yang membawakan fungsi manajemen risiko'), LOWER('dir. Bidang/sevp (unit pengguna)')) THEN 'TPP-2'
					WHEN LOWER(p.KEWENANGAN_PENGADAAN) = LOWER('Rapat Direksi') THEN 'TPP-3'
					ELSE 'Not Found'
			   END AS Kewenangan
		FROM PENGADAAN p
		WHERE p.STATUS_PENGADAAN = 'Done'
	) df
	GROUP BY df.Kewenangan
	ORDER BY df.Kewenangan ASC`
	
	if err := dashboardRepositoryImpl.DB.Raw(query).Scan(dashboardModel).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table((?) subquery, subQuery).Select(Kewenangan, COUNT(*) as Count).Group(Kewenangan).Scan(dashboardModel).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error {
	query := `
			SELECT df.STATUS_PENGADAAN, 
				COUNT(CASE WHEN df.jenis_pengadaan = 'IT' THEN 1 END) AS total_it,
				COUNT(CASE WHEN df.jenis_pengadaan not in ('IT','Premises') THEN 1 END) AS total_nonit,
				COUNT(CASE WHEN df.jenis_pengadaan = 'PREMISES' THEN 1 END) AS total_premises
			FROM (SELECT p.* FROM PENGADAAN p) df
			GROUP BY df.STATUS_PENGADAAN
			ORDER BY df.STATUS_PENGADAAN ASC`

	if err := dashboardRepositoryImpl.DB.Raw(query).Scan(statusPengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.STATUS_PENGADAAN, COUNT(*) as count_pengadaan).Group(p.STATUS_PENGADAAN ).Find(statusPengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	query := `SELECT df.METODE, 
		COUNT(CASE WHEN df.jenis_pengadaan = 'IT' THEN 1 END) AS total_it,
		COUNT(CASE WHEN df.jenis_pengadaan not in ('IT','Premises') THEN 1 END) AS total_nonit,
		COUNT(CASE WHEN df.jenis_pengadaan = 'PREMISES' THEN 1 END) AS total_premises
		FROM (SELECT p.* FROM PENGADAAN p WHERE p.STATUS_PENGADAAN = 'Done') df
		GROUP BY df.METODE
		ORDER BY df.METODE ASC`
	if err := dashboardRepositoryImpl.DB.Raw(query).Scan(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.METODE ,count(*) as count_metode).Group(p.METODE).Where(p.STATUS_PENGADAAN = ?,Done).Find(metodePengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneTrenPengadaanMasuk(c *fiber.Ctx,year string,metodePengadaan *[]map[string]interface{}) error {
	query := fmt.Sprintf(`
		WITH months AS (
			SELECT LEVEL AS month
			FROM DUAL
			CONNECT BY LEVEL <= 12
			),
			pengadaan_monthly AS (
			SELECT
			EXTRACT(MONTH FROM CAST(TO_TIMESTAMP(p.SCHEDULE_END_DATE, 'YYYY-MM-DD HH24:MI:SS.FF7') AS DATE)) AS MONTH
		FROM
			PENGADAAN p
		WHERE
			REGEXP_LIKE(p.SCHEDULE_END_DATE, '^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{7}$')
			AND EXTRACT(YEAR FROM CAST(TO_TIMESTAMP(p.SCHEDULE_END_DATE, 'YYYY-MM-DD HH24:MI:SS.FF7') AS DATE)) = %v
			AND p.STATUS = 'On Progress'
			)
			SELECT 
			m.month,
			COALESCE(COUNT(pm.month), 0) AS record_count
			FROM months m
			LEFT JOIN pengadaan_monthly pm ON m.month = pm.month
			GROUP BY m.month
		ORDER BY m.month
	`,year)
	if err := dashboardRepositoryImpl.DB.Raw(query).Scan(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.STATUS_PENGADAAN AS name,count(*) AS counting_data).Group(p.STATUS_PENGADAAN).Find(metodePengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneTrenPengadaanSelesai(c *fiber.Ctx,year string,metodePengadaan *[]map[string]interface{}) error {
	query := fmt.Sprintf(`
		WITH months AS (
			SELECT LEVEL AS month
			FROM DUAL
			CONNECT BY LEVEL <= 12
			),
			pengadaan_monthly AS (
			SELECT
			EXTRACT(MONTH FROM CAST(TO_TIMESTAMP(p.SCHEDULE_END_DATE, 'YYYY-MM-DD HH24:MI:SS.FF7') AS DATE)) AS MONTH
		FROM
			PENGADAAN_FILTER p
		WHERE
			REGEXP_LIKE(p.SCHEDULE_END_DATE, '^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{7}$')
			AND EXTRACT(YEAR FROM CAST(TO_TIMESTAMP(p.SCHEDULE_END_DATE, 'YYYY-MM-DD HH24:MI:SS.FF7') AS DATE)) = %v
			AND p.STATUS = 'Done'
			)
			SELECT 
			m.month,
			COALESCE(COUNT(pm.month), 0) AS record_count
			FROM months m
			LEFT JOIN pengadaan_monthly pm ON m.month = pm.month
			GROUP BY m.month
		ORDER BY m.month
	`,year)
	if err := dashboardRepositoryImpl.DB.Raw(query).Scan(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.STATUS_PENGADAAN AS name,count(*) AS counting_data).Group(p.STATUS_PENGADAAN).Find(metodePengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) InformasiRekanan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("DATA_VENDOR_RESULT dvr").Select("dvr.vendor_status_name,count(*) AS count_status_name").Group("dvr.vendor_status_name").Find(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(DATA_VENDOR_RESULT dvr).Select(dvr.vendor_status_name,count(*) AS count_status_name).Group(dvr.vendor_status_name).Find(metodePengadaan).Error; err != nil ")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) DataInformasiRekanan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Table("DATA_VENDOR_RESULT dvr").Select("dvr.NAME,DVR.vendor_activity_status_name").Where(`DVR.vendor_activity_status_name in ('Vendor Inaktif', 'Vendor Aktif')`).Find(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(DATA_VENDOR_RESULT dvr).Select(dvr.NAME,DVR.vendor_activity_status_name).Where(`DVR.vendor_activity_status_name in ('Vendor Inaktif', 'Vendor Aktif')`).Find(metodePengadaan).Error")
		return err
	}
	return nil
}