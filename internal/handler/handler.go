package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"invidious/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	routes := fiber.New()

	routes.Use(
		logger.New(),
		cors.New(),
	)

	routes.Get("/api", h.helloWorld)

	return routes
}
