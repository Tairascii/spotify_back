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

type TokenClaims struct {
	jwt.StandardClaims
	UserId    int  `json:"user_id"`
	IsRefresh bool `json:"is_refresh"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func NewAuthManager(repo repository.Auth) *AuthManager {
	return &AuthManager{repo: repo}
}

func (a *AuthManager) SignInUser(login, password string) (Tokens, error) {
	user, err := a.repo.GetUser(login, generatePassword(password))

	if err != nil {
		return Tokens{}, err
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		false,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(36 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		true,
	})

	singedAccess, err := accessToken.SignedString([]byte(signInKey))
	singedRefresh, err := refreshToken.SignedString([]byte(signInKey))
	return Tokens{AccessToken: singedAccess, RefreshToken: singedRefresh}, err
}

func (a *AuthManager) SignUpUser(user models.User) (int, error) {
	user.Password = generatePassword(user.Password)
	return a.repo.SignUpUser(user)
}

func (a *AuthManager) ParseToken(inputToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(inputToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}

		return []byte(signInKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok {
		return nil, errors.New("token claims invalid")
	}

	return claims, nil
}

func (a *AuthManager) RefreshTokens(refreshToken string) (Tokens, error) {
	claims, err := a.ParseToken(refreshToken)

	if err != nil {
		return Tokens{}, err
	}

	if !claims.IsRefresh {
		return Tokens{}, errors.New("invalid refresh token")
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		claims.UserId,
		false,
	})

	singedAccess, err := newAccessToken.SignedString([]byte(signInKey))
	return Tokens{AccessToken: singedAccess, RefreshToken: refreshToken}, nil
}

func generatePassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
