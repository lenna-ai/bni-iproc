package appconfig

import (
	"github.com/lenna-ai/bni-iproc/config"
	detailpengadaancontroller "github.com/lenna-ai/bni-iproc/controllers/detailPengadaanController"
	detailpengadaanrepositories "github.com/lenna-ai/bni-iproc/repositories/detailPengadaanRepositories"
	detailpengadaanservices "github.com/lenna-ai/bni-iproc/services/detailPengadaanServices"
)

func InitControllerServiceRepository() *detailpengadaancontroller.PengadaanControllerImpl {
	db := config.DB
	DetailPengadaanFilterRepository := detailpengadaanrepositories.NewDetailPengadaanRepository(db)
	DetailPengadaanFilteService := detailpengadaanservices.NewDetailPengadaanService(DetailPengadaanFilterRepository)
	DetailPengadaanFilteController := detailpengadaancontroller.NewDetailPengadaanController(DetailPengadaanFilteService)
	return DetailPengadaanFilteController
}
