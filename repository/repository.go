package repository

import (
	"github.com/jmoiron/sqlx"
	"spotify_back/repository/daos"
)

type Auth interface {
	SignInUser(login, password string) (string, error)
	SignUpUser(login, password string) (string, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Auth: daos.NewUserDao(db)}
}
