package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type pingController struct{}
type pingControllerInterface interface {
	Ping(c *fiber.Ctx) error
}

var (
	PingController pingControllerInterface = &pingController{}
)

func (ctr *pingController) Ping(c *fiber.Ctx) error {
	c.Status(http.StatusOK).SendString("ping ping pong")
	return nil
}
