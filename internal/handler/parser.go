package handler

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"invidious/internal/model"
	"time"
)

func (h *Handler) InsertChannels(c *fiber.Ctx) error {
	var input *model.Channels

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.CreatedAt = time.Now()
	err := h.repos.InsertIntoChannelsTable(context.Background(), input)

	return err
}

func (h *Handler) InsertPlaylists(c *fiber.Ctx) error {
	var input *model.Playlists

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.CreatedAt = time.Now()
	err := h.repos.InsertIntoPlaylistsTable(context.Background(), input)

	return err
}

func (h *Handler) InsertVideoMetas(c *fiber.Ctx) error {
	var input *model.VideoMetas

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.CreatedAt = time.Now()
	err := h.repos.InsertIntoVideoMetasTable(context.Background(), input)

	return err
}

func (h *Handler) InsertVideos(c *fiber.Ctx) error {
	var input *model.Videos

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.CreatedAt = time.Now()
	err := h.repos.InsertIntoVideosTable(context.Background(), input)

	return err
}

func (h *Handler) InsertGenres(c *fiber.Ctx) error {
	var input *model.Genres

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	err := h.repos.InsertIntoGenresTable(context.Background(), input)

	return err
}
