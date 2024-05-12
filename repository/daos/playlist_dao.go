package daos

import (
	"github.com/jmoiron/sqlx"
	"spotify_back/models"
)

type PlaylistDao struct {
	c *sqlx.DB
}

func NewPlaylistDao(c *sqlx.DB) *PlaylistDao {
	return &PlaylistDao{c: c}
}

func (dao *PlaylistDao) GetPlaylist(id int) (models.Playlist, error) {
	return models.Playlist{}, nil
}

func (dao *PlaylistDao) CreatePlaylist(playlist models.Playlist) (int, error) {
	return 0, nil
}

func (dao *PlaylistDao) DeletePlaylist(id int) error {
	return nil
}

func (dao *PlaylistDao) EditPlaylist(playlist models.Playlist) error {
	return nil
}

func (dao *PlaylistDao) AddSongToPlaylist(songId, playlistId int) error {
	return nil
}

func (dao *PlaylistDao) DeleteSongFromPlaylist(songId, playlistId int) error {
	return nil
}
