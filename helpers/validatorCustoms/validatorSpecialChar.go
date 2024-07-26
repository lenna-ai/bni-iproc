package validatorcustoms

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var CustomSpecialChar = "“<&>%$"

func CustomValidatorSpecialChar(fl validator.FieldLevel) bool {
    text := fl.Field().String()


    if strings.ContainsAny(text, CustomSpecialChar) {
        return false
    }

    return true
}