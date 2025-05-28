-- name: CreateUser: one
INSERT INTO users (
    name,
    email,
    password
)
VALUES (?, ?, ?);

-- name: GetUserById: one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByEmail: one
SELECT * FROM users
WHERE email = ?;

--name: DeleteUser: exec
DELETE FROM users
WHERE id = ?;