package services

import (
	"context"
	"errors"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/store/mysqlstore"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo utils.UserRepository
}

func NewUserService(repo utils.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func (us *UserService) CreateUser(ctx context.Context, name, email, password string) (entity.User, error) {

	hash, err := HashPassword(password)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	return us.repo.CreateUser(ctx, user)
}

func (us *UserService) AuthenticateUser(ctx context.Context, email, password string) (entity.User, string, error) {

	user, err := us.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return entity.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return entity.User{}, "", mysqlstore.ErrInvalidCredentials
		}
		return entity.User{}, "", err
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return entity.User{}, "", err
	}

	return user, token, nil

}
