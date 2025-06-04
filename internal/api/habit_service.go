package api

import (
	"context"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

type Service struct {
	repo HabitRepository
}

func NewService(repo HabitRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateHabit(ctx context.Context, habit entity.Habit) (int64, error) {
	return s.repo.CreateHabit(ctx, habit)
}

func (s *Service) ListHabits(ctx context.Context) ([]entity.Habit, error) {
	return s.repo.ListHabits(ctx)
}

func (s *Service) DeleteHabit(ctx context.Context, id int64) (bool, error) {
	return s.repo.DeleteHabit(ctx, id)
}
