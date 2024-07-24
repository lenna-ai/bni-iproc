package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	validatorcustoms "github.com/lenna-ai/bni-iproc/helpers/validatorCustoms"
)

func ResultUnauthorizedJsonApi(c *fiber.Ctx, data any, errorMessage string) error {
	return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
		"data":   data,
		"status": errorMessage,
	})
}

func ResultSuccessJsonApi(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusAccepted).JSON(data)
}

func ResultFailedJsonApi(c *fiber.Ctx, data any, errorMessage string) error {
	return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
		"data":   data,
		"status": errorMessage,
	})
}

func RecoverPanicContext(c *fiber.Ctx) error {
	if r := recover(); r != nil {
		err := fmt.Sprintf("Error occured %s", r)
		return ResultFailedJsonApi(c, nil, err)
	}
	return nil
}

func MessageErrorValidation(c *fiber.Ctx, field string, validateTag string, params string) error {
	var errorMessage string
	var messageValidator = map[string]any{
		"CustomValidatorSpecialChar": fmt.Sprintf("Please dont use char => %v",validatorcustoms.CustomSpecialChar),
	}

	if params != "" {
		errorMessage = fmt.Sprintf("The %v field/key is %v=%v. ", field, validateTag, params)
		
	} else {
		errorMessage = fmt.Sprintf("The %v field/key is %v. ", field, validateTag)
	}
	if validateTag != "" {
		for k, v := range messageValidator {
			if validateTag == k {
				errorMessage += fmt.Sprint(v)
			}
		}
	}
	return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
		"data":   nil,
		"status": errorMessage,
	})
}
