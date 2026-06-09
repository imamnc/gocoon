package user

import (
	"gocoon/core/middleware"

	"github.com/gofiber/fiber/v2"
)

func Mount(app *fiber.App) {

	app.Get("/user", middleware.AuthCheck, GetUser)
	app.Post("/user", RegisterUser)
	app.Patch("/user", UpdateUser)
	app.Delete("/user/:id", DeleteUser)

}
