package services

import (
	"context"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
)

type HabitService struct {
	repo utils.HabitRepository
}

func NewService(repo utils.HabitRepository) *HabitService {
	return &HabitService{repo: repo}
}

func (s *HabitService) CreateHabit(ctx context.Context, habit entity.Habit) (int64, error) {
	return s.repo.CreateHabit(ctx, habit)
}

func (s *HabitService) ListHabits(ctx context.Context) ([]entity.Habit, error) {
	return s.repo.ListHabits(ctx)
}

func (s *HabitService) DeleteHabit(ctx context.Context, id int64) (bool, error) {
	return s.repo.DeleteHabit(ctx, id)
}
