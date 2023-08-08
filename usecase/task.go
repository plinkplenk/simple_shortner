package usecase

import (
	"context"
	"github.com/plinkplenk/simple_shortner/domain"
)

type urlUsecase struct {
	urlRepository domain.UrlRepository
}

func NewUrlUsecase(c context.Context, urlRepository domain.UrlRepository) domain.UrlUsecase {
	return &urlUsecase{
		urlRepository: urlRepository,
	}
}

func (uu *urlUsecase) Create(c context.Context, url *domain.URL) error {
	return uu.urlRepository.Create(c, url)
}

func (uu *urlUsecase) GetByID(c context.Context, id string) (domain.URL, error) {
	return uu.urlRepository.GetByID(c, id)
}
