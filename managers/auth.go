package managers

import (
	"crypto/md5"
	"fmt"
	"spotify_back/models"
	"spotify_back/repository"
)

const salt = "y0so6aki9an4ru70"

type AuthManager struct {
	repo repository.Auth
}

func NewAuthManager(repo repository.Auth) *AuthManager {
	return &AuthManager{repo: repo}
}

func (a *AuthManager) SignInUser(login, password string) (string, error) {
	return "", nil
}

func (a *AuthManager) SignUpUser(user models.User) (int, error) {
	user.Password = generatePassword(user.Password)
	return a.repo.SignUpUser(user)
}

func generatePassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
