package handlers

import (
	"context"

	"demo-backend/use_cases/usermgmtuc"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type RegisterUseCase interface {
	Register(context.Context, usermgmtuc.RegisterRequest) (*usermgmtuc.RegisterResponse, error)
}

func RegisterHandler(useCase RegisterUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ctx = c.UserContext()
		var request = usermgmtuc.RegisterRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return errors.Wrap(err, "unable to parse incoming request")
		}

		response, err := useCase.Register(ctx, request)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}
