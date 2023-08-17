package dtos

import "github.com/gofrs/uuid"

type UserCreateDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDto struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password string  `json:"password"`
}

type UserResponseDto struct {
	ID       uuid.UUID `json:"ID"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
