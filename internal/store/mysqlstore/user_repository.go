package mysqlstore

import (
	"context"

	"github.com/google/uuid"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
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
		return entity.User{}, err
	}

	return user, nil

}
