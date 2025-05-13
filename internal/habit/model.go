package habit

import (
	"context"
	"time"
)

type Habit struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type HabitRepository interface {
	CreateHabit(ctx context.Context, name string) error
}
