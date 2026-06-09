package main

import (
	"fmt"

	"gocoon/core/config"
	"gocoon/core/database"
	"gocoon/core/models"
	"gocoon/services/auth"
	"gocoon/services/todo"
	"gocoon/services/user"
	"gocoon/services/welcome"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Prepare application
	Prepare()

	// Init fiber app
	app := fiber.New()

	// Setup routes
	SetupRoutes(app)

	// Expose app
	app.Listen(fmt.Sprintf(":%v", config.Data.Port))
}

func SetupRoutes(app *fiber.App) {
	// Welcome routes
	welcome.Mount(app)
	// User routes
	user.Mount(app)
	// Todo routes
	todo.Mount(app)
	// Auth routes
	auth.Mount(app)
}

func Prepare() {
	// Load env
	config.LoadEnv()
	// Load Config
	config.Load()
	// Load Models
	models.Register()
	// Connect to database
	database.Connect()
	// Migrate
	database.Migrate()
}
