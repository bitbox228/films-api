package service

import (
	vkfilms "github.com/bitbox228/vk-films-api"
	"github.com/bitbox228/vk-films-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user vkfilms.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, vkfilms.UserRole, error)
}

type Actor interface {
}

type Film interface {
}

type Service struct {
	Authorization
	Actor
	Film
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}