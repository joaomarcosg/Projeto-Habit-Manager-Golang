-- name: CreateHabit :execresult
INSERT INTO habits (
    name,
    category,
    description,
    frequency,
    start_date,
    target_date,
    priority
)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: ListHabits :many
SELECT * FROM habits
ORDER BY id;

-- name: GetHabitById :one
SELECT * FROM habits
WHERE id = ?;

-- name: DeleteHabit :exec
DELETE FROM habits
WHERE id = ?;

-- name: UpdateHabit :exec
UPDATE habits
SET
    name = ?,
    category = ?,
    description = ?,
    frequency = ?,
    start_date = ?,
    target_date = ?,
    priority = ?
WHERE id = ?;