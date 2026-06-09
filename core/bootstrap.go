package app

import (
	"fmt"

	"gocoon/config"
	"gocoon/database"
	"gocoon/models"

	"github.com/gofiber/fiber/v2"
)

func Init() {
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
	// Init fiber app
	app := fiber.New()
	// Init routes
	RegisterRoute(app)
	// Expose app
	app.Listen(fmt.Sprintf(":%v", config.Data.Port))
}
