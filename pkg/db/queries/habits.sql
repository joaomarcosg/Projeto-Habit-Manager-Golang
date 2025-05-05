-- name: ListHabits :many
SELECT * FROM habits;

-- name: CreateHabit :execresult
INSERT INTO habits(id, name, created_at) VALUES (?, ?, ?);