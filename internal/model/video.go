package model

import (
	"time"
)

type MediaInfo struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type (
	Channels struct {
		Id                     int         `json:"id"`
		ChannelTitle           string      `json:"author"`
		ChannelSlug            string      `json:"channelSlug"`
		ChannelDescription     string      `json:"description"`
		ChannelKeywords        []string    `json:"channelKeywords"`
		ChannelOwner           string      `json:"channelOwner"`
		ChannelYoutubeId       string      `json:"authorId"`
		ChannelBanners         []MediaInfo `json:"authorBanners"`
		ChannelThumbnails      []MediaInfo `json:"authorThumbnails"`
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
		PlaylistTitle       string    `json:"title"`
		PlaylistDescription string    `json:"playlistDescription"`
		PlaylistKeywords    []string  `json:"playlistKeywords"`
		PlaylistYoutubeId   string    `json:"playlistId"`
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
		VideoTitle           string    `json:"title"`
		VideoDescription     string    `json:"description"`
		VideoKeywords        []string  `json:"videoKeywords"`
		VideoYoutubeId       string    `json:"videoId"`
		VideoDownloadStatus  int       `json:"videoDownloadStatus"`
		VideoTranscodeStatus int       `json:"videoTranscodeStatus"`
		VideoLikes           int       `json:"videoLikes"`
		VideoDislikes        int       `json:"videoDislikes"`
		VideoViews           int       `json:"videoViews"`
		YoutubeLikes         int       `json:"likeCount"`
		YoutubeDislikes      int       `json:"dislikeCount"`
		YoutubeViews         int       `json:"viewCount"`
		CreatedAt            time.Time `json:"createdAt"`
		UpdatedAt            time.Time `json:"updatedAt"`
		DeletedAt            time.Time `json:"deletedAt"`
		VideoTranscodePath   string    `json:"videoTranscodePath"`
		VideoAccessType      int       `json:"videoAccessType"`
		GenreId              int       `json:"genreId"`
		Genre                string    `json:"genre"`
	}

	Genres struct {
		Id           int    `json:"id"`
		GenreTitle   string `json:"genreTitle"`
		GenreTitleRu string `json:"genreTitleRu"`
	}
)

type UniqId struct {
	UniqId string `json:"id"`
}

type Video struct {
	Elements []Videos `json:"videos"`
}

type PlaylistOfChannel struct {
	Playlists []Playlists `json:"playlists"`
}
