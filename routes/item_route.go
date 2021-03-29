package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/controllers"
)

func ItemRoutes(app *fiber.App) {
	app.Post("/api/items", controllers.ItemsController.Create)
	app.Post("/api/items/search", controllers.ItemsController.Search)
	app.Get("/api/items/:id", controllers.ItemsController.Get)
	app.Put("/api/items/:id", controllers.ItemsController.Update)
	app.Delete("/api/items/:id", controllers.ItemsController.Delete)
}
