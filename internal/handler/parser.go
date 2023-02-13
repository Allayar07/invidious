package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) helloWorld(c *fiber.Ctx) error {
	return c.SendString("hello World")
}
