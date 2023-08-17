package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plinkplenk/simple_shortner/internal/api/handlers"
)

func NewRegisterRouter(rh *handlers.RegisterHandler, router *fiber.Router) {
	//(*router).Get("/url/:id", rh.GetByID)
	(*router).Post("/users", rh.Create)
	//(*router).Get("/url/user/:user_id", rh.GetAllByUserID)
}
