package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"invidious/internal/repository"
)

type Handler struct {
	repos *repository.VideoRepository
}

func NewHandler(repos *repository.VideoRepository) *Handler {
	return &Handler{
		repos: repos,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	routes := fiber.New()

	routes.Use(
		logger.New(),
		cors.New(),
	)

	routes.Post("/channel", h.InsertChannels)
	routes.Post("/playlist", h.InsertPlaylists)
	routes.Post("/metas", h.InsertVideoMetas)
	routes.Post("/video", h.InsertVideos)
	routes.Post("/genres", h.InsertGenres)

	return routes
}
