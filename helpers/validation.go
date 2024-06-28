package helpers

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/lenna-ai/bni-iproc/config"
)

func ValidationFields(data any) (string, string, string, error) {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Pastikan data adalah struct setelah dereference pointer
	if v.Kind() != reflect.Struct {
		return "", "", "", errors.New("data bukan struct atau pointer ke struct")
	}

	// Validasi struct menggunakan validator
	err := config.Validate.Struct(data)
	if err != nil {
		for _, valueError := range err.(validator.ValidationErrors) {
			field, _ := v.Type().FieldByName(valueError.StructField())
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}
			return jsonTag, valueError.Tag(), valueError.Param(), errors.New("field catch validation")
		}
	}
	return "", "", "", nil
}
