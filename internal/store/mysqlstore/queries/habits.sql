-- name: CreateHabit :execresult
INSERT INTO habits (
    name,
    category,
    description,
    frequency,
    start_date,
    target_date,
    priority,
    user_id
)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: ListHabits :many
SELECT * FROM habits
WHERE user_id = ?
ORDER BY id;

-- name: GetHabitById :one
SELECT * FROM habits
WHERE id = ? AND user_id = ?;

-- name: DeleteHabit :exec
DELETE FROM habits
WHERE id = ? AND user_id = ?;

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
WHERE id = ? AND user_id = ?;