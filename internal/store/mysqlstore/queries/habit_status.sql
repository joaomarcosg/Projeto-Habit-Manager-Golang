-- name: UpdateHabitStatus :exec
INSERT INTO habit_status (habit_id, user_id, status, date)
VALUES (?, ?, ?, ?)
ON DUPLICATE KEY UPDATE
    status = VALUES(status),
    updated_at = CURRENT_TIMESTAMP;

-- name: HabitTrack :one
SELECT COUNT(*) FROM habit_status
WHERE habit_id = ? AND user_id = ? AND status = 'done'
AND date BETWEEN ? AND ?;