package api

import (
	"github.com/gofiber/fiber/v2"
)

type healthzApi struct {
}

func NewHelloApi() Interface {
	return &healthzApi{}
}

func (h *healthzApi) Register(app *fiber.App) {
	api := app.Group("/healthz")
	api.Get("", h.healthz)
}

func (h *healthzApi) healthz(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
