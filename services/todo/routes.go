package todo

import (
	"github.com/gofiber/fiber/v2"
)

func Mount(app *fiber.App) {
	app.Get("/todo", GetTodo)
	app.Post("/todo", CreateTodo)
	app.Patch("/todo", UpdateTodo)
	app.Delete("/todo/:id", DeleteTodo)
}
