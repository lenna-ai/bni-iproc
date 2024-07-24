package appconfig

import (
	"github.com/lenna-ai/bni-iproc/config"
	validatorcustoms "github.com/lenna-ai/bni-iproc/helpers/validatorCustoms"
)

func initValidator()  {
	config.Validate.RegisterValidation("CustomValidatorSpecialChar", validatorcustoms.CustomValidatorSpecialChar)
}

