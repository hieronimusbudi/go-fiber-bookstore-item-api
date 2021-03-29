package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/datasources/elasticsearch"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/datasources/mongodb"
	"github.com/hieronimusbudi/go-fiber-bookstore-item-api/routes"
)

func main() {
	app := fiber.New()
	routes.PingRoutes(app)
	routes.ItemRoutes(app)

	elasticsearch.Init()
	mongodb.Init()
	app.Listen(":9000")
}
