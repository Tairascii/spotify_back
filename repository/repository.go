package repository

import (
	"github.com/jmoiron/sqlx"
	"spotify_back/models"
	"spotify_back/repository/daos"
)

type Auth interface {
	GetUser(login, password string) (models.User, error)
	SignUpUser(user models.User) (int, error)
}

type Playlist interface {
	GetPlaylist(id int) (models.Playlist, error)
	CreatePlaylist(playlist models.Playlist) (int, error)
	DeletePlaylist(id int) error
	EditPlaylist(playlist models.Playlist) error
	AddSongToPlaylist(songId, playlistId int) error
	DeleteSongFromPlaylist(songId, playlistId int) error
}

type Repository struct {
	Auth
	Playlist
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Auth: daos.NewUserDao(db), Playlist: daos.NewPlaylistDao(db)}
}
