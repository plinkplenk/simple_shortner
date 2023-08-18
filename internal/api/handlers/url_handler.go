package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/plinkplenk/simple_shortner/internal/api/dtos"
	"github.com/plinkplenk/simple_shortner/internal/domain"
	"github.com/plinkplenk/simple_shortner/internal/utils"
	"net/http"
	"time"
)

type UrlHandler struct {
	UrlUsecase domain.URLUsecase
	Validator  *utils.UrlValidator
}

type map_ = fiber.Map

const (
	idLen              = 6
	defaultURLLifetime = 24 * time.Hour
)

func (uh *UrlHandler) Create(c *fiber.Ctx) error {
	expireIN := time.Now().Add(defaultURLLifetime).Unix()
	var input dtos.URLCreateDto
	urlID := utils.GenerateID(idLen)
	existingUrl, _ := uh.UrlUsecase.GetByID(c.Context(), urlID)
	for existingUrl != nil {
		urlID = utils.GenerateID(idLen)
		existingUrl, _ = uh.UrlUsecase.GetByID(c.Context(), urlID)
	}
	err := c.BodyParser(&input)
	if err != nil {
		return &Err{
			Code:    http.StatusBadRequest,
			Message: "field 'original' is required",
		}
	}
	original := input.Original
	if err := uh.Validator.Validate(original); err != nil {
		return &Err{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if t := input.Expire; t != nil {
		expireIN = *t
	}
	URLToCreate := domain.URL{
		ID:       urlID,
		Original: original,
		Expire:   expireIN,
	}
	newURL, err := uh.UrlUsecase.Create(c.Context(), &URLToCreate)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(
		map_{
			"url": dtos.URLResponseDto{
				ID:       newURL.ID,
				Original: newURL.Original,
				UserID:   newURL.UserID,
			},
		},
	)
}

func (uh *UrlHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	url, err := uh.UrlUsecase.GetByID(c.Context(), id)
	if err != nil {
		return &Err{
			Code:    http.StatusNotFound,
			Message: "URL with this ID not found",
		}
	}

	return c.JSON(
		map_{
			"url": dtos.URLResponseDto{
				ID:       url.ID,
				Original: url.Original,
				UserID:   url.UserID,
				Expire:   url.Expire,
			},
		},
	)
}

func (uh *UrlHandler) GetAllByUserID(c *fiber.Ctx) error {
	id := c.Params("user_id")
	userUUID, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	urls, err := uh.UrlUsecase.GetAllByUserID(c.Context(), userUUID)
	if err != nil {
		return err
	}
	urlResponse := make([]dtos.URLResponseDto, len(urls))
	for i, url := range urls {
		urlResponse[i] = dtos.URLResponseDto{
			ID:       url.ID,
			Original: url.Original,
			UserID:   url.UserID,
			Expire:   url.Expire,
		}
	}
	return c.JSON(map_{"urls": urlResponse})
}
