package repository

import (
	"context"
	"github.com/piaohao/godis"
	"github.com/plinkplenk/simple_shortner/domain"
)

type urlRepository struct {
	db         *godis.Redis
	collection string
}

func NewURLRepository(db *godis.Redis, collection string) domain.UrlRepository {
	return &urlRepository{
		db:         db,
		collection: collection,
	}
}

func (ur *urlRepository) Create(c context.Context, url *domain.URL) error {
	_, err := ur.db.Set(url.ID, url.Original)
	return err
}

func (ur *urlRepository) GetByID(c context.Context, id string) (domain.URL, error) {
	url := domain.URL{}
	original, err := ur.db.Get(id)
	if err != nil {
		return url, err
	}
	url.ID = id
	url.Original = original
	return url, nil
}
