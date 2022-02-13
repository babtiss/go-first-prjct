package service

import (
	"crypto/sha1"
	"fmt"
	todo "go-application/model"
	"go-application/pkg/repository"
)

const HashCode = "1frercfr"

type AuthService struct {
	Repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{Repo: repo}
}

func (s AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generateHashForPassword(user.Password)
	return s.Repo.CreateUser(user)
}

func generateHashForPassword(password string) string {

	hashForPassword := sha1.New()
	hashForPassword.Write([]byte(password))

	return fmt.Sprintf("%x", hashForPassword.Sum([]byte(HashCode)))
}
