package managers

import (
	"mime/multipart"
	"spotify_back/models"
	"spotify_back/repository"
)

type Auth interface {
	SignInUser(login, password string) (Tokens, error)
	SignUpUser(user models.User) (int, error)
	ParseToken(token string) (*TokenClaims, error)
	RefreshTokens(refreshToken string) (Tokens, error)
}

type Playlist interface {
	GetPlaylist(playlistId int) (PlaylistWithSongs, error)
	CreatePlaylist(userId int) (int, error)
	DeletePlaylist(id int) error
	EditPlaylist(playlist models.Playlist) error
	AddSongToPlaylist(songId, playlistId int) error
	DeleteSongFromPlaylist(songId, playlistId int) error
}

type Song interface {
	UploadFileSong(file *multipart.File) (string, error)
}

type Manager struct {
	Auth
	Playlist
	Song
}

func NewManager(repo *repository.Repository) *Manager {
	return &Manager{
		Auth:     NewAuthManager(repo.Auth),
		Playlist: NewPlaylistManager(repo.Playlist),
		Song:     NewSongManager(repo.Song),
	}
}
