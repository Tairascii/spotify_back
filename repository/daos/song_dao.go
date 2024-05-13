package daos

import "github.com/jmoiron/sqlx"

type SongDao struct {
	c *sqlx.DB
}
