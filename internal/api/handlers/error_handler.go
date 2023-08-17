package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type defaultError struct {
	Message string `json:"message"`
}
type errorResponse struct {
	Err defaultError `json:"error"`
}

type Err = fiber.Error

func APIErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	errBody := defaultError{Message: err.Error()}
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(code).JSON(errorResponse{errBody})
}
