package monitoringcontroller

import (
	"encoding/json"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/bni-iproc/config"
	"github.com/lenna-ai/bni-iproc/helpers"
	gormhelpers "github.com/lenna-ai/bni-iproc/helpers/gormHelpers"
	jwthelpers "github.com/lenna-ai/bni-iproc/helpers/jwtHelpers"
	"github.com/lenna-ai/bni-iproc/models/prosesPengadaanModel/formatters"
)

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) JenisPengadaan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)
	jenisPengadaan, err := monitoringProsesPengadaanImpl.MonitoringProsesPengadaan.JenisPengadaan(c)
	if err != nil {
		log.Printf("error monitoringProsesPengadaanImpl.MonitoringProsesPengadaan.JenisPengadaan %v \n ", err)
		return helpers.ResultFailedJsonApi(c, jenisPengadaan, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c, jenisPengadaan)
}

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) DetailProsesPengadaan(c *fiber.Ctx) error {
	defer helpers.RecoverPanicContext(c)

	var totalCount = new(int64)
	jenisPengadaanParams,_:=url.QueryUnescape(c.Params("jenis_pengadaan"))
	jenis_pengadaan := strings.ToLower(jenisPengadaanParams)
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	getPengadaanFormatter, err := monitoringProsesPengadaanImpl.MonitoringProsesPengadaan.DetailProsesPengadaan(c,totalCount,jenis_pengadaan)
	if err != nil {
		log.Printf("error monitoringProsesPengadaanImpl.MonitoringProsesPengadaan.getPengadaanFormatter %v \n ", err)
		return helpers.ResultFailedJsonApi(c, getPengadaanFormatter, err.Error())
	}
	return helpers.ResultSuccessJsonApi(c,gormhelpers.PaginatedResponse(page,pageSize,*totalCount,getPengadaanFormatter))
}

func (monitoringProsesPengadaanImpl *MonitoringProsesPengadaanImpl) PutProsesPengadaan(c *fiber.Ctx) error {

	defer helpers.RecoverPanicContext(c)

	putPengadaanFormatter := new(formatters.PutPengadaanFormatter)
	if err := c.BodyParser(putPengadaanFormatter); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}
	err := config.Validate.Struct(putPengadaanFormatter)
	if err != nil {
		for _, valueError := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(*putPengadaanFormatter).FieldByName(valueError.StructField())
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}
			return helpers.MessageErrorValidation(c, jsonTag, valueError.Tag(), valueError.Param())
		}
	}
	var dataMe = new(map[string]any)
	jwthelpers.MeJwt(c,dataMe)
	data,_ := json.Marshal(dataMe)
	datame := string(data)
	putPengadaanFormatter.DELETED_BY = &datame
	*putPengadaanFormatter.DELETED_BY = string(data)
	if err := monitoringProsesPengadaanImpl.MonitoringProsesPengadaan.PutProsesPengadaan(c, putPengadaanFormatter); err != nil {
		return helpers.ResultFailedJsonApi(c, nil, err.Error())
	}

	return helpers.ResultSuccessJsonApi(c, putPengadaanFormatter)
	
}
