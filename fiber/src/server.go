package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		context.SendStatus(400)
		return context.JSON(fiber.Map{"name": "Max Mustermann"})
	})

	app.Get("/products", func(context *fiber.Ctx) error {
		products := []string{"Milk", "Bread", "Honey", "Potatoes", "Chocolate"}
		return context.JSON(products)
	})

	app.Listen(":3000")
}
