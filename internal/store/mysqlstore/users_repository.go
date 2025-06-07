package mysqlstore

import (
	"context"
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

var ErrDuplicatedEmailOrUsername = errors.New("name or email already exists")

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
