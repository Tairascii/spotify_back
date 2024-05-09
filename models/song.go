package models

import "time"

type Song struct {
	Id        int       `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	SongPath  string    `json:"song_path"`
	ImagePath string    `json:"image_path"`
	Author    string    `json:"author"`
	UserId    int       `json:"user_id"`
}
