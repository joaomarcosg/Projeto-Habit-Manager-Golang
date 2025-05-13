package habit

import "context"

type Service struct {
	repo HabitRepository
}

func NewService(repo HabitRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateHabit(ctx context.Context, name string) error {
	return s.repo.CreateHabit(ctx, name)
}
