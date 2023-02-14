package model

import (
	"time"
)

type (
	Channels struct {
		Id                     int         `json:"id"`
		ChannelTitle           string      `json:"channelTitle"`
		ChannelSlug            string      `json:"channelSlug"`
		ChannelDescription     string      `json:"channelDescription"`
		ChannelKeywords        []string    `json:"channelKeywords"`
		ChannelOwner           string      `json:"channelOwner"`
		ChannelYoutubeId       string      `json:"channelYoutubeId"`
		ChannelBanners         interface{} `json:"channel_banners"`
		ChannelThumbnails      interface{} `json:"channel_thumbnails"`
		ChannelException       string      `json:"channelException"`
		ChannelStatus          int         `json:"channelStatus"`
		ChannelIsForeign       bool        `json:"channelIsForeign"`
		ChannelSubscriberCount int         `json:"channelSubscriberCount"`
		CreatedAt              time.Time   `json:"createdAt"`
		UpdatedAt              time.Time   `json:"updatedAt"`
		DeletedAt              time.Time   `json:"deletedAt"`
		ChannelAccessType      int         `json:"channelAccessType"`
	}

	Playlists struct {
		Id                  int       `json:"id"`
		ChannelId           int       `json:"channelId"`
		PlaylistTitle       string    `json:"playlistTitle"`
		PlaylistDescription string    `json:"playlistDescription"`
		PlaylistKeywords    []string  `json:"playlistKeywords"`
		PlaylistYoutubeId   string    `json:"playlistYoutubeId"`
		PlaylistVideoCount  int       `json:"playlistVideoCount"`
		CreatedAt           time.Time `json:"createdAt"`
		UpdatedAt           time.Time `json:"updatedAt"`
		DeletedAt           time.Time `json:"deletedAt"`
	}

	VideoMetas struct {
		Id                      int         `json:"id"`
		VideoId                 int         `json:"videoId"`
		VideoDuration           int         `json:"videoDuration"`
		VideoThumbnails         interface{} `json:"video-thumbnails"`
		CreatedAt               time.Time   `json:"createdAt"`
		UpdatedAt               time.Time   `json:"updatedAt"`
		DeletedAt               time.Time   `json:"deletedAt"`
		ImageFullDownloadStatus int         `json:"imageFullDownloadStatus"`
	}

	Videos struct {
		Id                   int       `json:"id"`
		ChannelId            int       `json:"channelId"`
		PlaylistId           int       `json:"playlistId"`
		VideoTitle           string    `json:"videoTitle"`
		VideoDescription     string    `json:"videoDescription"`
		VideoKeywords        []string  `json:"videoKeywords"`
		VideoYoutubeId       string    `json:"videoYoutubeId"`
		VideoDownloadStatus  int       `json:"videoDownloadStatus"`
		VideoTranscodeStatus int       `json:"videoTranscodeStatus"`
		VideoLikes           int       `json:"videoLikes"`
		VideoDislikes        int       `json:"videoDislikes"`
		VideoViews           int       `json:"videoViews"`
		YoutubeLikes         int       `json:"youtubeLikes"`
		YoutubeDislikes      int       `json:"youtubeDislikes"`
		YoutubeViews         int       `json:"youtubeViews"`
		CreatedAt            time.Time `json:"createdAt"`
		UpdatedAt            time.Time `json:"updatedAt"`
		DeletedAt            time.Time `json:"deletedAt"`
		VideoTranscodePath   string    `json:"videoTranscodePath"`
		VideoAccessType      int       `json:"videoAccessType"`
		GenreId              int       `json:"genreId"`
	}

	Genres struct {
		Id           int    `json:"id"`
		GenreTitle   string `json:"genreTitle"`
		GenreTitleRu string `json:"genreTitleRu"`
	}
)
