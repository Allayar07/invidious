package repository

import (
	"context"
	"fmt"
	"invidious/internal/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	ChannelsTable   = "channels"
	PlaylistsTable  = "playlists"
	VideoMetasTable = "video_metas"
	VideosTable     = "videos"
	GenresTable     = "genres"
)

type VideoRepository struct {
	db *pgxpool.Pool
}

func NewVideoRepository(db *pgxpool.Pool) *VideoRepository {
	return &VideoRepository{
		db: db,
	}
}

func (r *VideoRepository) InsertIntoChannelsTable(ctx context.Context, channels model.Channels) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (channel_title, channel_slug, channel_description, channel_keywords,
                channel_owner, channel_youtube_id, channel_banners, channel_thumbnails, channel_exception, channel_status, 
                channel_is_foreign, channel_subscriber_count, created_at, updated_at, deleted_at, channel_access_type) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`, ChannelsTable)

	row := r.db.QueryRow(ctx, query, channels.ChannelTitle, channels.ChannelSlug, channels.ChannelDescription, channels.ChannelKeywords,
		channels.ChannelOwner, channels.ChannelYoutubeId, channels.ChannelBanners, channels.ChannelThumbnails, channels.ChannelException,
		channels.ChannelStatus, channels.ChannelIsForeign, channels.ChannelSubscriberCount, channels.CreatedAt, channels.UpdatedAt, channels.DeletedAt,
		channels.ChannelAccessType)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *VideoRepository) InsertIntoPlaylistsTable(ctx context.Context, playlist model.Playlists) error {
	query := fmt.Sprintf(`INSERT INTO %s (channel_id, playlist_title, playlist_description, playlist_keywords,
                				playlist_youtube_id, playlist_video_count, created_at, updated_at, deleted_at) 
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, PlaylistsTable)

	_, err := r.db.Exec(ctx, query, playlist.ChannelId, playlist.PlaylistTitle, playlist.PlaylistDescription, playlist.PlaylistKeywords,
		playlist.PlaylistYoutubeId, playlist.PlaylistVideoCount, playlist.CreatedAt, playlist.UpdatedAt, playlist.DeletedAt)

	return err
}

func (r *VideoRepository) InsertIntoVideoMetasTable(ctx context.Context, videoMeta model.VideoMetas) error {
	query := fmt.Sprintf("INSERT INTO %s (video_id, video_duration, video_thumbnails, created_at, updated_at, deleted_at, image_full_download_status) VALUES ($1, $2, $3, $4, $5, $6, $7)", VideoMetasTable)

	_, err := r.db.Exec(ctx, query, videoMeta.VideoId, videoMeta.VideoDuration, videoMeta.VideoThumbnails, videoMeta.CreatedAt, videoMeta.UpdatedAt, videoMeta.DeletedAt, videoMeta.ImageFullDownloadStatus)

	return err
}

func (r *VideoRepository) InsertIntoVideosTable(ctx context.Context, video model.Videos) error {
	query := fmt.Sprintf(`INSERT INTO %s (channel_id, playlist_id, video_title, video_description, video_keywords,
                video_youtube_id, video_download_status, video_transcode_status, video_likes, video_dislikes, video_views, 
                youtube_likes, youtube_dislikes, youtube_views, created_at, updated_at, deleted_at, video_transcode_path, video_access_type, genre_id)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`, VideosTable)

	_, err := r.db.Exec(ctx, query, video.ChannelId, video.PlaylistId, video.VideoTitle, video.VideoDescription,
		video.VideoKeywords, video.VideoYoutubeId, video.VideoDownloadStatus, video.VideoTranscodeStatus, video.VideoLikes,
		video.VideoDislikes, video.VideoViews, video.YoutubeLikes, video.YoutubeDislikes, video.YoutubeViews, video.CreatedAt,
		video.UpdatedAt, video.DeletedAt, video.VideoTranscodePath, video.VideoAccessType, video.GenreId)

	return err
}

func (r *VideoRepository) InsertIntoGenresTable(ctx context.Context, genres model.Genres) error {
	query := fmt.Sprintf(`INSERT INTO %s (genre_title, genre_title_ru) VALUES ($1, $2)`, GenresTable)

	_, err := r.db.Exec(ctx, query, genres.GenreTitle, genres.GenreTitleRu)

	return err
}
