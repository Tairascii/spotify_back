package managers

import (
	"spotify_back/models"
	"spotify_back/repository"
)

type Auth interface {
	SignInUser(login, password string) (string, error)
	SignUpUser(user models.User) (int, error)
	ParseToken(token string) (int, error)
}

type Manager struct {
	Auth
}

func NewManager(repo *repository.Repository) *Manager {
	return &Manager{Auth: NewAuthManager(repo.Auth)}
}
