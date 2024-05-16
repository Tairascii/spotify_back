package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"spotify_back/models"
	"strings"
)

type PlaylistSongDao struct {
	c *sqlx.DB
}

func NewPlaylistSongDao(c *sqlx.DB) *PlaylistSongDao {
	return &PlaylistSongDao{c: c}
}

func (dao *PlaylistSongDao) CreateLink(playlistId, songId int) (int, error) {
	var id int
	query := fmt.Sprintf("insert into playlist_song (playlist_id, song_id) values ($1 $2) returning id")
	row := dao.c.QueryRow(query, playlistId, songId)

	if err := row.Scan(&id); err != nil {
		log.Printf("something went wrong while creating link %s", err.Error())
		return 0, err
	}

	return id, nil
}

func (dao *PlaylistSongDao) GetSongsByPlaylist(playlistId int) ([]models.Song, error) {
	query := fmt.Sprintf("select song_id from playlist_song where playlist_id=$1")
	rows, err := dao.c.Query(query, playlistId)
	defer rows.Close()

	if err != nil {
		log.Printf("something went wrong while getting song ids %s", err.Error())
		return nil, err
	}

	songIds := make([]int, 10)

	for rows.Next() {
		var songId int
		err = rows.Scan(&songId)
		if err != nil {
			log.Printf("something went wrong while scanning song ids %s", err.Error())
		}
		songIds = append(songIds, songId)
	}

	placeholders := make([]string, len(songIds))
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	query = fmt.Sprintf("select * from songs where id in (%s)", strings.Join(placeholders, ","))
	songRows, err := dao.c.Query(query, songIds)

	if err != nil {
		log.Printf("something went wrong while getting songs %s", err.Error())
		return nil, err
	}

	songs := make([]models.Song, 10)

	for songRows.Next() {
		var song models.Song
		err = songRows.Scan(&song)
		if err != nil {
			log.Printf("something went wrong while scanning song ids %s", err.Error())
		}
		songs = append(songs, song)
	}
	return songs, nil
}
