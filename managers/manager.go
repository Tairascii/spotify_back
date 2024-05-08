package managers

import "spotify_back/repository"

type Manager struct {
	repo *repository.Repository
}

func NewManager(repo *repository.Repository) *Manager {
	return &Manager{repo}
}
