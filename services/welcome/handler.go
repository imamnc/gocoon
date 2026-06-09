package welcome

import (
	"github.com/gofiber/fiber/v2"
)

// This is a handler to welcoming the user at main routes
func Welcome(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"author":  "Imam Nc",
		"message": "Welcome to Gocoon Boilerplate, REST API boilerplate built with Go, Fiber, and GORM.",
	})
}
