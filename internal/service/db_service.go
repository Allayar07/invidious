package service

import (
	"context"
	"encoding/json"
	"fmt"
	"invidious/internal/model"
	"invidious/internal/repository"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

type DbService struct {
	repos *repository.VideoRepository
}

func NewDbService(repo *repository.VideoRepository) *DbService {
	return &DbService{
		repos: repo,
	}
}

func (s *DbService) InsertIntoChannelsTable(ctx context.Context, id model.UniqId) (int, error) {

	channels := model.Channels{}

	urlOfChannel := fmt.Sprintf("https://vid.puffyan.us/api/v1/channels/%s?fields=author,authorId,authorBanners,authorThumbnails,description", id.UniqId)
	_, respBody, err := fasthttp.Get(nil, urlOfChannel)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(respBody, &channels); err != nil {
		return 0, err
	}

	mediaUrls := make([]model.MediaInfo, 0)
	thumbnailUrls := make([]model.MediaInfo, 0)

	banner2560 := ChangeSpecifiedUrl("kjlkh", channels.ChannelBanners, 2560)
	thumbnail48 := ChangeSpecifiedUrl("jhjhd", channels.ChannelThumbnails, 48)
	thumbnail512 := ChangeSpecifiedUrl("jhjhd", channels.ChannelThumbnails, 512)

	mediaUrls = append(mediaUrls, banner2560)
	thumbnailUrls = append(thumbnailUrls, thumbnail48, thumbnail512)
	channels.ChannelBanners = mediaUrls
	channels.CreatedAt = time.Now()
	channels.UpdatedAt = time.Now()
	channels.ChannelKeywords = []string{}
	channels.ChannelThumbnails = thumbnailUrls

	chanId, err := s.repos.InsertIntoChannelsTable(ctx, channels)

	return chanId, err
}

func (s *DbService) InsertIntoGenresTable(ctx context.Context, genre model.Genres) error {
	return s.repos.InsertIntoGenresTable(ctx, genre)
}

func (s *DbService) InsertIntoPlaylistsTable(ctx context.Context, id model.UniqId) ([]model.Playlists, error) {

	playlist := model.PlaylistOfChannel{}

	urlOfChannelsPlaylists := fmt.Sprintf("https://vid.puffyan.us/api/v1/channels/%s/playlists", id.UniqId)

	_, respBody, err := fasthttp.Get(nil, urlOfChannelsPlaylists)
	if err != nil {
		return []model.Playlists{}, err
	}

	if err = json.Unmarshal(respBody, &playlist); err != nil {
		return []model.Playlists{}, err
	}

	playlistarray := make([]model.Playlists, 0)
	for _, obj := range playlist.Playlists {
		obj.PlaylistKeywords = []string{}
		obj.ChannelId = 1

		if err := s.repos.InsertIntoPlaylistsTable(ctx, obj); err != nil {
			return []model.Playlists{}, err
		}

		playlistarray = append(playlistarray, obj)
	}

	return playlistarray, nil
}

func (s *DbService) InsertIntoVideosTable(ctx context.Context, id model.UniqId) ([]model.Videos, error) {
	var (
		videosArray model.Video
		video       model.Videos
	)
	urlOfVideos := fmt.Sprintf("https://vid.puffyan.us/api/v1/channels/%s/videos", id.UniqId)

	_, respBody, err := fasthttp.Get(nil, urlOfVideos)
	if err != nil {
		return []model.Videos{}, err
	}

	if err = json.Unmarshal(respBody, &videosArray); err != nil {
		return []model.Videos{}, err
	}

	videoarray := make([]model.Videos, 0)

	var wg sync.WaitGroup
	for _, obj := range videosArray.Elements {

		wg.Add(1)

		go func(obj model.Videos) error {
			urlOfVideoGenre := fmt.Sprintf("https://vid.puffyan.us/api/v1/videos/%s", obj.VideoYoutubeId)

			_, videorespBody, err := fasthttp.Get(nil, urlOfVideoGenre)
			if err != nil {
				return err
			}

			if err = json.Unmarshal(videorespBody, &video); err != nil {
				return err
			}

			switch video.Genre {
			case "Autos & Vehicles":
				video.GenreId = 1
			case "Comedy":
				video.GenreId = 2
			case "Education":
				video.GenreId = 3
			case "Entertainment":
				video.GenreId = 4
			case "Gaming":
				video.GenreId = 5
			case "Howto & Style":
				video.GenreId = 6
			case "Music":
				video.GenreId = 7
			case "News & Politics":
				video.GenreId = 8
			case "Nonprofits & Activism":
				video.GenreId = 9
			case "Pets & Animals":
				video.GenreId = 10
			case "Science & Technology":
				video.GenreId = 11
			case "Sports":
				video.GenreId = 12
			case "Travel & Events":
				video.GenreId = 13
			case "People & Blogs":
				video.GenreId = 14
			case "Film & Animation":
				video.GenreId = 15
			}

			video.VideoKeywords = []string{}
			video.PlaylistId = 1
			video.ChannelId = 1
			video.CreatedAt = time.Now()
			video.UpdatedAt = time.Now()

			videoarray = append(videoarray, video)

			err = s.repos.InsertIntoVideosTable(ctx, video)

			wg.Done()

			return err
		}(obj)

	}

	wg.Wait()

	return videoarray, nil
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
