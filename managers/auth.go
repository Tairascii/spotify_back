package managers

import (
	"crypto/md5"
	"errors"
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

func (a *AuthManager) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}

		return []byte(signInKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claims invalid")
	}

	return claims.UserId, nil
}

func generatePassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
