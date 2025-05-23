package mysqlstore

import (
	"context"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

type HabitRepository struct {
	q *Queries
}

func NewHabitRepository(q *Queries) *HabitRepository {
	return &HabitRepository{q: q}
}

func (r *HabitRepository) CreateHabit(ctx context.Context, habit entity.Habit) (int64, error) {
	result, err := r.q.CreateHabit(ctx, CreateHabitParams{
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency: NullHabitsFrequency{
			HabitsFrequency: HabitsFrequency(habit.Frequency),
			Valid:           true,
		},
		StartDate:  habit.StartDate,
		TargetDate: habit.TargetDate,
		Priority:   habit.Priority,
	})

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
