package usecase

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/plinkplenk/simple_shortner/internal/domain"
	"time"
)

type urlUsecase struct {
	urlRepository domain.URLRepository
	timeout       time.Duration
}

func NewUrlUsecase(urlRepository domain.URLRepository, timeout time.Duration) domain.URLUsecase {
	return &urlUsecase{
		urlRepository: urlRepository,
		timeout:       timeout,
	}
}

func (uu *urlUsecase) Create(c context.Context, url *domain.URL) (*domain.URL, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()
	return uu.urlRepository.Create(ctx, url)
}

func (uu *urlUsecase) GetByID(c context.Context, id string) (*domain.URL, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()
	return uu.urlRepository.GetByID(ctx, id)
}

func (uu *urlUsecase) GetAllByUserID(c context.Context, id uuid.UUID) ([]*domain.URL, error) {
	ctx, cancel := context.WithTimeout(c, uu.timeout)
	defer cancel()
	return uu.urlRepository.GetAllByUserID(ctx, id)
}
