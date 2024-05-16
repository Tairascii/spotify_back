package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"spotify_back/models"
)

type SongDao struct {
	c *sqlx.DB
}

func NewSongDao(c *sqlx.DB) *SongDao {
	return &SongDao{c: c}
}

func (dao *SongDao) CreateSong(song models.Song) (int, error) {
	var id int
	query := fmt.Sprintf("insert into songs (created_at, title, song_path, image_path, author, user_id) values ($1, $2, $3, $4, $5, $6) returning id")
	row := dao.c.QueryRow(query, song.CreatedAt, song.Title, song.SongPath, song.ImagePath, song.Author, song.UserId)

	if err := row.Scan(&id); err != nil {
		log.Printf("something went wrong while creating song %s", err.Error())
		return 0, err
	}
	return id, nil
}

func (dao *SongDao) DeleteSong(songId int) error {
	query := fmt.Sprintf("delete from songs where id = $1")
	_, err := dao.c.Exec(query, songId)

	return err
}

func (dao *SongDao) GetSong(songId int) (models.Song, error) {
	var song models.Song
	query := fmt.Sprintf("select * from songs where id = $1")
	err := dao.c.Get(&song, query, songId)

	if err != nil {
		log.Printf("something went wrong while getting song %s", err.Error())
		return models.Song{}, err
	}

	return song, nil
}
