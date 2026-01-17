package controllers

import (
	"adaptivetesting/src/application/dto/requests"
	"adaptivetesting/src/application/usecases"
	webapi "adaptivetesting/src/infrastructure/http"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserControllers struct {
	app      *fiber.App
	useCases *usecases.User
}

func (this *UserControllers) RegisterUserController(c *fiber.Ctx) error {
	request, err := webapi.ValidateAndGetBody[requests.RegisterUserRequest](c)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := this.useCases.Register(ctx, *request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}

func (this *UserControllers) AuthUserController(c *fiber.Ctx) error {
	request, err := webapi.ValidateAndGetBody[requests.AuthUserRequest](c)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := this.useCases.Authorize(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
