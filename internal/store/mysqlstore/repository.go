package mysqlstore

import (
	"context"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/habit"
	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/store/mysqlstore"
)

func toNullHabitsFrequency(freq string) mysqlstore.NullHabitsFrequency {
	if freq == "" {
		return mysqlstore.NullHabitsFrequency{
			Valid: false,
		}
	}

	return mysqlstore.NullHabitsFrequency{
		HabitsFrequency: mysqlstore.HabitsFrequency(freq),
		Valid:           true,
	}
}

type HabitRepository struct {
	q *Queries
}

func NewHabitRepository(q *Queries) *HabitRepository {
	return &HabitRepository{q: q}
}

func (r *HabitRepository) CreateHabit(ctx context.Context, habit habit.Habit) (int64, error) {
	result, err := r.q.CreateHabit(ctx, CreateHabitParams{
		Name:        habit.Name,
		Category:    habit.Category,
		Description: habit.Description,
		Frequency:   toNullHabitsFrequency(habit.Frequency),
		StartDate:   habit.StartDate,
		TargetDate:  habit.TargetDate,
		Priority:    habit.Priority,
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
