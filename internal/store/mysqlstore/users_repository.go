package mysqlstore

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

var (
	ErrDuplicatedEmailOrUsername = errors.New("name or email already exists")
	ErrInvalidCredentials        = errors.New("invalid credentials")
)

type UserRepository struct {
	q *Queries
}

func NewUserRepository(q *Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {

	user.ID = uuid.New().String()

	_, err := r.q.CreateUser(ctx, CreateUserParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			if strings.Contains(mysqlErr.Message, "email") {
				return entity.User{}, ErrDuplicatedEmailOrUsername
			}
			if strings.Contains(mysqlErr.Message, "name") {
				return entity.User{}, ErrDuplicatedEmailOrUsername
			}
			return entity.User{}, ErrDuplicatedEmailOrUsername
		}
		return entity.User{}, err
	}

	return user, nil

}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {

	user, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, ErrInvalidCredentials
		}
		return entity.User{}, err
	}

	return entity.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil

}
