package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"invidious/internal/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func (h *Handler) InsertChannels(c *fiber.Ctx) error {
	var (
		input   model.ChannelId
		channel model.Channels
	)

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	urlOfChannel := fmt.Sprintf("https://vid.puffyan.us/api/v1/channels/%s", input.ChannelUniqId)

	_, respBody, err := fasthttp.Get(nil, urlOfChannel)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, &channel); err != nil {
		return err
	}

	mediaUrls := make([]model.MediaInfo, 0)
	thumbnailUrls := make([]model.MediaInfo, 0)

	banner2560 := ChangeSpecifiedUrl("kjlkh", channel.ChannelBanners, 2560)
	thumbnail48 := ChangeSpecifiedUrl("jhjhd", channel.ChannelThumbnails, 48)
	thumbnail512 := ChangeSpecifiedUrl("jhjhd", channel.ChannelThumbnails, 512)

	mediaUrls = append(mediaUrls, banner2560)
	thumbnailUrls = append(thumbnailUrls, thumbnail48, thumbnail512)
	//thumbnailUrl[1] = ChangeSpecifiedUrl("shfafs", thumbnailUrl, 1, 48, 48)
	//thumbnailUrl[5] = ChangeSpecifiedUrl("jhgjh", thumbnailUrl, 5, 512, 512)
	channel.ChannelBanners = mediaUrls
	channel.CreatedAt = time.Now()
	channel.UpdatedAt = time.Now()
	channel.ChannelKeywords = []string{}
	channel.ChannelThumbnails = thumbnailUrls

	fmt.Println(channel.ChannelBanners)
	err = h.repos.InsertIntoChannelsTable(context.Background(), channel)
	if err != nil {
		return err
	}

	return c.JSON(channel)
}

func (h *Handler) InsertPlaylists(c *fiber.Ctx) error {
	var input model.Playlists

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.CreatedAt = time.Now()
	err := h.repos.InsertIntoPlaylistsTable(context.Background(), input)

	return err
}

func (h *Handler) InsertVideoMetas(c *fiber.Ctx) error {
	var input model.VideoMetas

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.CreatedAt = time.Now()
	err := h.repos.InsertIntoVideoMetasTable(context.Background(), input)

	return err
}

func (h *Handler) InsertVideos(c *fiber.Ctx) error {
	var input model.Videos

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	input.CreatedAt = time.Now()
	err := h.repos.InsertIntoVideosTable(context.Background(), input)

	return err
}

func (h *Handler) InsertGenres(c *fiber.Ctx) error {
	var input model.Genres

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	err := h.repos.InsertIntoGenresTable(context.Background(), input)

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
