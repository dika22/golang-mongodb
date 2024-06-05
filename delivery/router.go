package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func routerGroup(app *fiber.App, handler handler) {
	app.Use(recover.New())

	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return nil })

	v1 := app.Group("/v1")
	{
		v1.Post("/order", handler.orderHandler.Order)
		v1.Get("/order/:id", handler.orderHandler.FindByID)
		v1.Get("/order/list", handler.orderHandler.List)
		v1.Put("/order/update", handler.orderHandler.Update)
	}
}