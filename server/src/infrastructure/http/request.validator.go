package webapi

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validatorSgt = validator.New()

func ValidateAndGetBody[RequestType any](c *fiber.Ctx) (*RequestType, error) {
	var request RequestType
	if err := c.BodyParser(&request); err != nil {
		return nil, err
	}

	if err := validatorSgt.Struct(&request); err != nil {
		return nil, err
	}

	return &request, nil
}
