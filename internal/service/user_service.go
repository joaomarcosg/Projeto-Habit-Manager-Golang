package service

import (
	"context"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/api"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo api.UserRepository
}

func NewUserService(repo api.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (s *UserService) CreateUser(ctx context.Context, name, email, password string) (entity.User, error) {

	hash, err := HashPassword(password)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	return s.repo.CreateUser(ctx, user)
}
