package service

import (
	todo "go-application/model"
	"go-application/pkg/repository"
)

type AuthService struct {
	Repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s AuthService) CreateUser(user todo.User) (int, error) {
	return s.Repo.CreateUser(user)
}
