-- name: CreateUser :one
INSERT INTO
    users (
        id,
        created_at,
        updated_at,
        email,
        hashed_password
    )
VALUES (
        gen_random_uuid (),
        NOW(),
        NOW(),
        $1,
        $2
    )
RETURNING
    *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByRefreshToken :one
SELECT users.id, users.created_at, users.updated_at, users.email
FROM users
    INNER JOIN refresh_tokens ON user_id = refresh_tokens.user_id
WHERE
    refresh_tokens.token = $1
    AND refresh_tokens.revoked_at IS NULL
    AND refresh_tokens.expires_at > NOW();

-- name: DeleteUsers :exec
DELETE FROM users;