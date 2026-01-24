-- name: CreateUser :exec
INSERT INTO users(id,username,ratings)
VALUES (
    gen_random_uuid(),
    $1,
    $2
);