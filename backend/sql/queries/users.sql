-- name: CreateUser :exec
INSERT INTO users(id,username,ratings)
VALUES (
    gen_random_uuid(),
    $1,
    $2
);

-- name: GetAllUsers :many
SELECT username,ratings
FROM users;