package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/plinkplenk/simple_shortner/internal/api/handlers"
	"github.com/plinkplenk/simple_shortner/internal/domain"
	"github.com/plinkplenk/simple_shortner/internal/repository"
	"github.com/plinkplenk/simple_shortner/internal/usecase"
	"github.com/plinkplenk/simple_shortner/internal/utils"
	"log"
	"time"
)

func Setup(app *fiber.App, pool *pgxpool.Pool) {
	urlRepository := repository.NewURLRepository(pool, domain.CollectionURL)
	userRepository := repository.NewUserRepository(pool, domain.CollectionUsers)
	urlValidator, err := utils.NewURLValidator()
	userValidator := utils.UserValidator{}
	if err != nil {
		log.Fatal(err)
	}
	uh := &handlers.UrlHandler{
		UrlUsecase: usecase.NewUrlUsecase(urlRepository, 1*time.Second),
		Validator:  urlValidator,
	}
	rh := &handlers.RegisterHandler{
		RegisterUsecase: usecase.NewRegisterUsecase(userRepository, 1*time.Second),
		Validator:       userValidator,
	}
	lh := &handlers.LoginHandler{
		LoginUsecase: usecase.NewLoginUsecase(userRepository, 1*time.Second),
		Validator:    userValidator,
	}

	api := app.Group("/api", logger.New())
	NewURLRouter(uh, &api)
	NewRegisterRouter(rh, &api)
	NewLoginRouter(lh, &api)
}
