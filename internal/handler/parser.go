package handler

import (
	"context"
	"fmt"
	"invidious/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) InsertChannels(c *fiber.Ctx) error {
	var input model.UniqId

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	channelId, err := h.service.DB.InsertIntoChannelsTable(context.Background(), input)
	if err != nil {
		return err
	}

	return c.JSON(channelId)
}

func (h *Handler) InsertVideos(c *fiber.Ctx) error {
	var input model.UniqId

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	video, err := h.service.DB.InsertIntoVideosTable(context.Background(), input)
	if err != nil {
		return err
	}

	return c.JSON(video)
}

func (h *Handler) InsertToPlaylist(c *fiber.Ctx) error {
	var input model.UniqId

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	playlists, err := h.service.DB.InsertIntoPlaylistsTable(context.Background(), input)
	if err != nil {
		return err
	}

	return c.JSON(playlists)
}

func (h *Handler) InsertGenres(c *fiber.Ctx) error {
	var input model.Genres

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	err := h.service.DB.InsertIntoGenresTable(context.Background(), input)

	return err
}

func ChangeSpecifiedUrl(filePath string, objects []model.MediaInfo, width int) model.MediaInfo {

	for _, obj := range objects {

		if obj.Width == width {
			obj.Url = fmt.Sprintf("assets/images/channel/thumbnails/%s", filePath)
			return obj
		}
	}

	return model.MediaInfo{}

}
