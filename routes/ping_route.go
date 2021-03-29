package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/controllers"
)

func PingRoutes(app *fiber.App) {
	app.Get("/api/items/ping", controllers.PingController.Ping)
}
