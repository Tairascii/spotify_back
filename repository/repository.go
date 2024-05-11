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

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Auth: daos.NewUserDao(db)}
}
