package mysqlstore

type HabitRepository struct {
	q *Queries
}

func NewHabitRepository(q *Queries) *HabitRepository {
	return &HabitRepository{q: q}
}
