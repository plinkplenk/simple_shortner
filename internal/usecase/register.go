package usecase

import (
	"context"
	"github.com/plinkplenk/simple_shortner/internal/domain"
	"time"
)

type registerUsecase struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewRegisterUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		timeout:        timeout,
	}
}

func (ru *registerUsecase) Create(c context.Context, user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, ru.timeout)
	defer cancel()
	return ru.userRepository.Create(ctx, user)
}
func (ru *registerUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, ru.timeout)
	defer cancel()
	return ru.userRepository.GetByEmail(ctx, email)
}

func (ru *registerUsecase) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, ru.timeout)
	defer cancel()
	return ru.userRepository.GetByUsername(ctx, username)
}
