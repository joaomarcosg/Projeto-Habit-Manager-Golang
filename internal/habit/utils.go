package habit

import "context"

type HabitRepository interface {
	CreateHabit(ctx context.Context, name string) error
}
