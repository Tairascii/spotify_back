package managers

import (
	"spotify_back/models"
	"spotify_back/repository"
)

type PlaylistManager struct {
	repo repository.Playlist
}

type PlaylistWithSongs struct {
	PlaylistInfo models.Playlist
	Songs        []models.Song
}

func NewPlaylistManager(repo repository.Playlist) *PlaylistManager {
	return &PlaylistManager{repo: repo}
}

func (man *PlaylistManager) CreatePlaylist(userId int) (int, error) {
	return 0, nil
}

func (man *PlaylistManager) GetPlaylist(playlistId int) (PlaylistWithSongs, error) {
	return PlaylistWithSongs{}, nil
}

func (man *PlaylistManager) EditPlaylist(playlist models.Playlist) error {
	return nil
}

func (man *PlaylistManager) DeletePlaylist(playlistId int) error {
	return nil
}

func (man *PlaylistManager) AddSongToPlaylist(songId, playlistId int) error {
	return nil
}

func (man *PlaylistManager) DeleteSongFromPlaylist(songId, playlistId int) error {
	return nil
}
