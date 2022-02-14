package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo "go-application/model"
	"go-application/pkg/repository"
	"time"
)

const (
	TokenLifeTime = 24 * time.Hour
	PasswordHash  = "1frerYcfr45fb"
	JWTSigned     = "fdgergYerf45t67h"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthService struct {
	Repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{Repo: repo}
}
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.Repo.GetUser(username, generateHashForPassword(password))
	if err != nil {
		return "", err
	}
	// 1 аргумент - стандартный метод подписи токена
	// 2 аргумент - json object с полями ( поставил время жизни токену 24 часа)
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(TokenLifeTime).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.Id,
		})

	return token.SignedString([]byte(JWTSigned))
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.PasswordHash = generateHashForPassword(user.PasswordHash)
	return s.Repo.CreateUser(user)
}

func (s *AuthService) ParseJWT(tokenString string) (int, error) {

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSigned), nil
	})

	if err != nil {
		return 0, errors.New("что-то сломалось")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("struct this token != struct our tokenClaims")
	}
	return claims.UserId, nil
}

func generateHashForPassword(password string) string {

	hashForPassword := sha1.New()
	hashForPassword.Write([]byte(password))

	return fmt.Sprintf("%x", hashForPassword.Sum([]byte(PasswordHash)))
}
