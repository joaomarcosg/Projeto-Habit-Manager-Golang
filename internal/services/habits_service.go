package services

import (
	"context"
	"time"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/utils"
)

type HabitService struct {
	repo utils.HabitRepository
}

func NewHabitService(repo utils.HabitRepository) *HabitService {
	return &HabitService{repo: repo}
}

func (s *HabitService) CreateHabit(ctx context.Context, userID string, habit entity.Habit) (int64, error) {
	return s.repo.CreateHabit(ctx, userID, habit)
}

func (s *HabitService) ListHabits(ctx context.Context, userID string) ([]entity.Habit, error) {
	return s.repo.ListHabits(ctx, userID)
}

func (s *HabitService) DeleteHabit(ctx context.Context, userID string, id int64) (bool, error) {
	return s.repo.DeleteHabit(ctx, userID, id)
}

func (s *HabitService) UpdateHabit(ctx context.Context, userID string, id int64, habit entity.Habit) (int64, error) {
	return s.repo.UpdateHabit(ctx, userID, id, habit)
}

func (s *HabitService) UpdateHabitStatus(ctx context.Context, userID string, id int64, status string, date time.Time) (bool, error) {
	return s.repo.UpdateHabitStatus(ctx, userID, id, status, date)
}
