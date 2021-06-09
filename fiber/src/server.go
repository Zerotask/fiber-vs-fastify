package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/zerotask/fiber-vs-fastify/src/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.Status(400).JSON(fiber.Map{"name": "Max Mustermann"})
	})

	app.Get("/products", func(context *fiber.Ctx) error {
		products := []string{"Milk", "Bread", "Honey", "Potatoes", "Chocolate"}
		return context.JSON(products)
	})

	routes.AddTodos(app)

	app.Listen(":3000")
}
