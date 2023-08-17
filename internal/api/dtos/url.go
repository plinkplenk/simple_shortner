package dtos

import (
	"github.com/gofrs/uuid"
)

type URLCreateDto struct {
	Original string `json:"original"`
	Expire   *int   `json:"expire"`
}

type URLResponseDto struct {
	ID       string     `json:"id"`
	Original string     `json:"original"`
	Expire   int        `json:"expire"`
	UserID   *uuid.UUID `json:"user_id"`
}
