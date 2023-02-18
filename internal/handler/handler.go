package handler

import (
	"invidious/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	routes.Post("/channel", h.InsertChannels)
	routes.Post("/playlist", h.InsertToPlaylist)
	// routes.Post("/metas", h.InsertVideoMetas)
	routes.Post("/video", h.InsertVideos)
	routes.Post("/genres", h.InsertGenres)

	return routes
}
