package managers

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"spotify_back/models"
	"spotify_back/repository"
	"time"
)

const (
	salt      = "y0so6aki9an4ru70"
	signInKey = "ya6ipr0kat1lt3bya"
)

type AuthManager struct {
	repo repository.Auth
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthManager(repo repository.Auth) *AuthManager {
	return &AuthManager{repo: repo}
}

func (a *AuthManager) SignInUser(login, password string) (string, error) {
	user, err := a.repo.GetUser(login, generatePassword(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	singed, err := token.SignedString([]byte(signInKey))
	return singed, err
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
