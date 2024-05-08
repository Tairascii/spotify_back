package managers

import "spotify_back/repository"

type Auth interface {
	SignInUser(login, password string) (string, error)
}

type Manager struct {
	Auth
}

func NewManager(repo *repository.Repository) *Manager {
	return &Manager{Auth: NewAuthManager(repo.Auth)}
}
