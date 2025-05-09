package mysqlstore

import "context"

type HabitRepository struct {
	q *Queries
}

func NewHabitRepository(q *Queries) *HabitRepository {
	return &HabitRepository{q: q}
}

func (r *HabitRepository) CreateHabit(ctx context.Context, name string) error {
	_, err := r.q.CreatHabit(ctx, name)
	return err
}
