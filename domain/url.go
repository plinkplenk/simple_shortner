package domain

import "context"

const CollectionUrl = "urls"

type URL struct {
	ID       string
	Original string
}

type UrlRepository interface {
	Create(c context.Context, url *URL) error
	GetByID(id string) (URL, error)
}

type UrlUsecase interface {
	Create(c context.Context, url *URL) error
	GetByID(c context.Context, id string) (URL, error)
}
