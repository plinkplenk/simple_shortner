package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plinkplenk/simple_shortner/internal/api/handlers"
)

func NewURLRouter(uh *handlers.UrlHandler, router *fiber.Router) {
	(*router).Get("/url/:id", uh.GetByID)
	(*router).Post("/url", uh.Create)
	(*router).Get("/url/user/:user_id", uh.GetAllByUserID)
}
