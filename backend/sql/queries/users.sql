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

-- name: GetRandomUsers :many
SELECT username
FROM users
ORDER BY RANDOM()
LIMIT $1;

-- name: BatchUpdateUserRating :exec
UPDATE users
SET ratings=updates.new_rating
FROM(
    SELECT
        UNNEST($1::text[]) AS username,
        UNNEST($2::int[]) AS new_rating
) AS updates
WHERE users.username=updates.username;

-- name: GetUsersByUsername :many
SELECT username,ratings
FROM users
WHERE username ILIKE '%' || $1 || '%'
AND username>$2
ORDER BY username
LIMIT $3;