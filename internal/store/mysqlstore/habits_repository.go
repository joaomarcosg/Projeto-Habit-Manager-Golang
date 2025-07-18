package mysqlstore

import (
	"context"
	"time"

	"github.com/joaomarcosg/Projeto-Habit-Manager-Golang/internal/entity"
)

type HabitRepository struct {
	q *Queries
}

func NewHabitRepository(q *Queries) *HabitRepository {
	return &HabitRepository{q: q}
}

func (r *HabitRepository) CreateHabit(ctx context.Context, userID string, habit entity.Habit) (int64, error) {
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
		UserID:     userID,
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

func (r *HabitRepository) DeleteHabit(ctx context.Context, userID string, id int64) (bool, error) {

	err := r.q.DeleteHabit(ctx, DeleteHabitParams{
		ID:     int32(id),
		UserID: userID,
	})
	if err != nil {
		return false, err
	}

	return true, nil

}

func (r *HabitRepository) UpdateHabit(ctx context.Context, userID string, id int64, habit entity.Habit) (int64, error) {
	err := r.q.UpdateHabit(ctx, UpdateHabitParams{
		ID:          int32(id),
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
		UserID:     userID,
	})

	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *HabitRepository) UpdateHabitStatus(ctx context.Context, userID string, id int64, status string, date time.Time) (bool, error) {

	err := r.q.UpdateHabitStatus(ctx, UpdateHabitStatusParams{
		HabitID: int32(id),
		UserID:  userID,
		Status:  HabitStatusStatus(status),
		Date:    date,
	})

	if err != nil {
		return false, err
	}

	return true, nil

}

func (r *HabitRepository) HabitTrack(ctx context.Context, userID string, id int64, startDate, targetDate time.Time) (int64, error) {

	track, err := r.q.HabitTrack(ctx, HabitTrackParams{
		HabitID:  int32(id),
		UserID:   userID,
		FromDate: startDate,
		ToDate:   targetDate,
	})

	if err != nil {
		return 0, err
	}

	return track, nil

}
