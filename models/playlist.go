package models

import "time"

type Playlist struct {
	Id          int       `json:"id"`
	CreateAt    time.Time `json:"create_at"`
	Author      string    `json:"author"`
	ImagePath   string    `json:"image_path"`
	UserId      int       `json:"user_id"`
	Description string    `json:"description"`
}

type PlaylistSong struct {
	PlaylistId int `json:"playlist_id"`
	SongId     int `json:"song_id"`
}
