package domain

import (
	"context"
	"github.com/gofrs/uuid"
)

const CollectionURL = "urls"

type URL struct {
	ID       string     `db:"id"`
	Original string     `db:"original"`
	Expire   int        `db:"expire"`
	UserID   *uuid.UUID `db:"user_id"`
}

type URLRepository interface {
	Create(ctx context.Context, url *URL) (*URL, error)
	GetByID(ctx context.Context, id string) (*URL, error)
	GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]*URL, error)
}

type URLUsecase interface {
	Create(ctx context.Context, url *URL) (*URL, error)
	GetByID(ctx context.Context, id string) (*URL, error)
	GetAllByUserID(ctx context.Context, userID uuid.UUID) ([]*URL, error)
}
