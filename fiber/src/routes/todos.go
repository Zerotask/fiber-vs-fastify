package routes

import "github.com/gofiber/fiber/v2"

func AddTodos(app *fiber.App) {
	todos := app.Group("/todos")

	todos.Get("/", func(context *fiber.Ctx) error {
		return context.Status(201).SendString("Todos get test")
	})
}
