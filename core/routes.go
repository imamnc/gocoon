package app

import (
	"gocoon/services/auth"
	"gocoon/services/todo"
	"gocoon/services/user"
	"gocoon/services/welcome"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(app *fiber.App) {
	// Welcome routes
	welcome.Mount(app)
	// User routes
	user.Mount(app)
	// Todo routes
	todo.Mount(app)
	// Auth routes
	auth.Mount(app)
}
