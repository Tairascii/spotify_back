package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type LikedSongDao struct {
	c *sqlx.DB
}

func NewLikedSongDao(c *sqlx.DB) *LikedSongDao {
	return &LikedSongDao{c: c}
}

func (dao *LikedSongDao) AddLike(userId, songId int, createdAt time.Time) (int, error) {
	var id int
	query := fmt.Sprintf("insert into liked_songs (user_id, song_id, created_at) values ($1, $2, $3) returning id")
	row := dao.c.QueryRow(query, userId, songId, createdAt)

	if err := row.Scan(&id); err != nil {
		log.Printf("something went wrong while adding like %s", err.Error())
		return 0, err
	}
	return id, nil
}

func (dao *LikedSongDao) RemoveLike(userId, songId int) error {
	query := fmt.Sprintf("delete from liked_songs where used_id = $1 and song_id = $2")
	row := dao.c.QueryRow(query, userId, songId)

	return row.Err()
}
