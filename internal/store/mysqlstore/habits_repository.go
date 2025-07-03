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

func (r *HabitRepository) ListHabits(ctx context.Context, userID string) ([]entity.Habit, error) {

	result, err := r.q.ListHabits(ctx, userID)
	if err != nil {
		return []entity.Habit{}, err
	}

	habits := make([]entity.Habit, len(result))

	for i, h := range result {
		freq := ""
		if h.Frequency.Valid {
			freq = string(h.Frequency.HabitsFrequency)
		}
		habits[i] = entity.Habit{
			ID:          int64(h.ID),
			Name:        h.Name,
			Category:    h.Category,
			Description: h.Description,
			Frequency:   freq,
			StartDate:   h.StartDate,
			TargetDate:  h.TargetDate,
			Priority:    h.Priority,
		}
	}

	return habits, nil

}

func (r *HabitRepository) DeleteHabit(ctx context.Context, id int64) (bool, error) {

	err := r.q.DeleteHabit(ctx, int32(id))
	if err != nil {
		return false, err
	}

	return true, nil

}
