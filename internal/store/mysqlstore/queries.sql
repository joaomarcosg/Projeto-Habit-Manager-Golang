-- name: CreatHabit :execresult
INSERT INTO habits (name)
VALUES (?);

-- name: ListHabits :many
SELECT * FROM habits
ORDER BY created_at DESC;