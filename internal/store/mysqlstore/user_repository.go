package mysqlstore

type UserRepository struct {
	q *Queries
}

func NewUserRepository(q *Queries) *UserRepository {
	return &UserRepository{q: q}
}
