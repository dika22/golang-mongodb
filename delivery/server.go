package delivery

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Http(handler handler) *fiber.App {
	app := fiber.New(
		fiber.Config{},
	)

	app.Use(cors.New())

	// init router
	routerGroup(app, handler)

	app.Use(func(c *fiber.Ctx) error {
		resp := fiber.Map{
			"status":  fmt.Sprintf("route %s or method not allowed", http.StatusText(http.StatusNotFound)),
			"message": fmt.Sprintf("route %s", http.StatusText(http.StatusNotFound)),
		}
		return c.Status(fiber.StatusNotFound).JSON(resp)
	})

	return app
}
