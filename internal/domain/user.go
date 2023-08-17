package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

const CollectionUsers = "users"

type User struct {
	ID       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*User, error)
}
