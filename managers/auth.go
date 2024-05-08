package managers

import "spotify_back/repository"

type AuthManager struct {
	repo repository.Auth
}

func NewAuthManager(repo repository.Auth) *AuthManager {
	return &AuthManager{repo: repo}
}

func (a *AuthManager) SignInUser(login, password string) (string, error) {
	return "", nil
}
