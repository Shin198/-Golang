-- name: CreateUser :execresult
INSERT INTO users (create_at, update_at, name)
VALUES (?,?,?)