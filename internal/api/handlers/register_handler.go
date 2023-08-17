package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plinkplenk/simple_shortner/internal/api/dtos"
	"github.com/plinkplenk/simple_shortner/internal/domain"
	"github.com/plinkplenk/simple_shortner/internal/utils"
	u "github.com/plinkplenk/simple_shortner/pkg/utils/password_hashing"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterHandler struct {
	RegisterUsecase domain.RegisterUsecase
	Validator       utils.UserValidator
}

func (rh *RegisterHandler) Create(c *fiber.Ctx) error {
	var input dtos.UserCreateDto
	err := c.BodyParser(&input)
	if err != nil {
		return err
	}
	if user, _ := rh.RegisterUsecase.GetUserByEmail(c.Context(), input.Email); user != nil {
		return &Err{
			Code:    http.StatusBadRequest,
			Message: "User with this email already exists",
		}
	}
	if user, _ := rh.RegisterUsecase.GetUserByUsername(c.Context(), input.Email); user != nil {
		return &Err{
			Code:    http.StatusBadRequest,
			Message: "User with this username already exists",
		}
	}
	if err := rh.Validator.ValidateUserData(&input); err != nil {
		return &Err{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	hashedPassword, err := u.HashPassword(input.Password, bcrypt.MinCost)
	if err != nil {
		return &Err{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	userToCreate := domain.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	newUser, err := rh.RegisterUsecase.Create(c.Context(), &userToCreate)
	if err != nil {
		return &Err{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(
		m{
			"user": dtos.UserResponseDto{
				ID:       newUser.ID,
				Username: newUser.Username,
				Email:    newUser.Email,
			},
		},
	)
}
