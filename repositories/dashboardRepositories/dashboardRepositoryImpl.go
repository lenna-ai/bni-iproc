package dashboardrepositories

import (
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
	subQuery := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select(`CASE
		WHEN p.KEWENANGAN_PENGADAAN = 'Pemimpin Departemen Divisi PFA (Unit Pelaksana)' OR p.KEWENANGAN_PENGADAAN = 'Pemimpin Departemen (Unit Pengguna)' THEN 'TPD-1'
		WHEN p.KEWENANGAN_PENGADAAN = 'Pemimpin Divisi PFA (Unit Pelaksana)' OR p.KEWENANGAN_PENGADAAN = 'Pemimpin Divisi/Satuan (Unit Pengguna)' THEN 'TPD-2'
		WHEN p.KEWENANGAN_PENGADAAN = 'Direktur yang membawahkan fungsi pengadaan' OR p.KEWENANGAN_PENGADAAN = 'Direktur yang membawakan fungsi manajemen risiko' OR p.KEWENANGAN_PENGADAAN = 'Dir. Bidang/SEVP (Unit Pengguna)' THEN 'TPP-1'
		WHEN p.KEWENANGAN_PENGADAAN = 'DIRUT' OR p.KEWENANGAN_PENGADAAN = 'WADIRUT' OR p.KEWENANGAN_PENGADAAN = 'Direktur yang membawahkan fungsi pengadaan' OR p.KEWENANGAN_PENGADAAN = 'Direktur yang membawakan fungsi manajemen risiko' OR p.KEWENANGAN_PENGADAAN = 'Dir. Bidang/SEVP (Unit Pengguna)' THEN 'TPP-2'
		WHEN p.KEWENANGAN_PENGADAAN = 'Rapat Direksi' THEN 'TPP-3'
		ELSE 'Not Found'
		END AS Kewenangan`).Where("p.STATUS_PENGADAAN = ?", "On Progress")
	
	if err := dashboardRepositoryImpl.DB.Table("(?) subquery", subQuery).Select("Kewenangan, COUNT(*) as Count").Group("Kewenangan").Scan(dashboardModel).Error; err != nil {
		log.Println("err := dashboardRepositoryImpl.DB.Table((?) subquery, subQuery).Select(Kewenangan, COUNT(*) as Count).Group(Kewenangan).Scan(dashboardModel).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select("p.STATUS_PENGADAAN, COUNT(*) as count_pengadaan").Group("p.STATUS_PENGADAAN ").Find(statusPengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.STATUS_PENGADAAN, COUNT(*) as count_pengadaan).Group(p.STATUS_PENGADAAN).Find(statusPengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select("p.METODE ,count(*) as count_metode").Group("p.METODE").Where("p.STATUS_PENGADAAN = ?","On Progress").Find(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.METODE ,count(*) as count_metode).Group(p.METODE).Where(p.STATUS_PENGADAAN = ?,On Progress).Find(metodePengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingKeputusan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select(`p.NAMA,p.JENIS_PENGADAAN, sum(p.NILAI_PENGADAAN_INISASI) as estimasi_nilai_pengadaan,CASE
	WHEN sum(p.NILAI_PENGADAAN_INISASI) < 500000000 THEN 'TPD1'
	WHEN sum(p.NILAI_PENGADAAN_INISASI) < 3000000000 THEN 'TPD2'
	WHEN sum(p.NILAI_PENGADAAN_INISASI) < 75000000000 THEN 'TPP1'
	WHEN sum(p.NILAI_PENGADAAN_INISASI) < 150000000000 THEN 'TPP2'
	WHEN sum(p.NILAI_PENGADAAN_INISASI) > 150000000000 THEN 'TPP3'
	END AS KEWENANGAN_PENGADAAN`).Group("p.JENIS_PENGADAAN,p.NAMA").Scan(metodePengadaan).Error; err != nil {
		log.Println("(dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnGoingKeputusan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneKewenangan(c *fiber.Ctx,dashboardModel *[]map[string]interface{}) error {
	subQuery := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select(`CASE
		WHEN p.KEWENANGAN_PENGADAAN = 'Pemimpin Departemen Divisi PFA (Unit Pelaksana)' OR p.KEWENANGAN_PENGADAAN = 'Pemimpin Departemen (Unit Pengguna)' THEN 'TPD-1'
		WHEN p.KEWENANGAN_PENGADAAN = 'Pemimpin Divisi PFA (Unit Pelaksana)' OR p.KEWENANGAN_PENGADAAN = 'Pemimpin Divisi/Satuan (Unit Pengguna)' THEN 'TPD-2'
		WHEN p.KEWENANGAN_PENGADAAN = 'Direktur yang membawahkan fungsi pengadaan' OR p.KEWENANGAN_PENGADAAN = 'Direktur yang membawakan fungsi manajemen risiko' OR p.KEWENANGAN_PENGADAAN = 'Dir. Bidang/SEVP (Unit Pengguna)' THEN 'TPP-1'
		WHEN p.KEWENANGAN_PENGADAAN = 'DIRUT' OR p.KEWENANGAN_PENGADAAN = 'WADIRUT' OR p.KEWENANGAN_PENGADAAN = 'Direktur yang membawahkan fungsi pengadaan' OR p.KEWENANGAN_PENGADAAN = 'Direktur yang membawakan fungsi manajemen risiko' OR p.KEWENANGAN_PENGADAAN = 'Dir. Bidang/SEVP (Unit Pengguna)' THEN 'TPP-2'
		WHEN p.KEWENANGAN_PENGADAAN = 'Rapat Direksi' THEN 'TPP-3'
		ELSE 'Not Found'
		END AS Kewenangan`).Where("p.STATUS_PENGADAAN = ?", "Done")
	
	if err := dashboardRepositoryImpl.DB.Table("(?) subquery", subQuery).Select("Kewenangan, COUNT(*) as Count").Group("Kewenangan").Scan(dashboardModel).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table((?) subquery, subQuery).Select(Kewenangan, COUNT(*) as Count).Group(Kewenangan).Scan(dashboardModel).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneStatus(c *fiber.Ctx,statusPengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select("p.STATUS_PENGADAAN, COUNT(*) as count_pengadaan").Group("p.STATUS_PENGADAAN ").Find(statusPengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.STATUS_PENGADAAN, COUNT(*) as count_pengadaan).Group(p.STATUS_PENGADAAN ).Find(statusPengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneMetode(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select("p.METODE ,count(*) as count_metode").Group("p.METODE").Where("p.STATUS_PENGADAAN = ?","Done").Find(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.METODE ,count(*) as count_metode).Group(p.METODE).Where(p.STATUS_PENGADAAN = ?,Done).Find(metodePengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) PengadaanOnDoneTrenPengadaan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Table("PENGADAAN p").Select("p.STATUS_PENGADAAN AS name,count(*) AS counting_data").Group("p.STATUS_PENGADAAN").Find(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(PENGADAAN p).Select(p.STATUS_PENGADAAN AS name,count(*) AS counting_data).Group(p.STATUS_PENGADAAN).Find(metodePengadaan).Error; err != nil")
		return err
	}
	return nil
}

func (dashboardRepositoryImpl *DashboardRepositoryImpl) InformasiRekanan(c *fiber.Ctx,metodePengadaan *[]map[string]interface{}) error {
	if err := dashboardRepositoryImpl.DB.Scopes(gormhelpers.Paginate(c)).Table("DATA_VENDOR_RESULT dvr").Select("dvr.vendor_activity_status_name,count(*) AS count_status_name").Group("dvr.vendor_activity_status_name").Find(metodePengadaan).Error; err != nil {
		log.Println("dashboardRepositoryImpl.DB.Table(DATA_VENDOR_RESULT dvr).Select(dvr.vendor_activity_status_name,count(*) AS count_status_name).Group(dvr.vendor_activity_status_name).Find(metodePengadaan).Error; err != nil ")
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