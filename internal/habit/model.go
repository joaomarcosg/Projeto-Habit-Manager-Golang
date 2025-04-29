package habit

import "time"

type Habit struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}
