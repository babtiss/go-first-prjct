package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	todo "go-application/model"
	"go-application/pkg/repository"
	"time"
)

const HashCode = "1frercfr"

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
	// 1 аргумент - стандартный метод (для подписей)
	// 2 аргумент - json object с полями ( поставил время жизни токену 24 часа)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	tokenCode := "fdgergerf45t67h"
	return token.SignedString([]byte(tokenCode))
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generateHashForPassword(user.Password)
	return s.Repo.CreateUser(user)
}

func generateHashForPassword(password string) string {

	hashForPassword := sha1.New()
	hashForPassword.Write([]byte(password))

	return fmt.Sprintf("%x", hashForPassword.Sum([]byte(HashCode)))
}
