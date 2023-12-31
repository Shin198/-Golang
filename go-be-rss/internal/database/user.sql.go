// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (create_at, update_at, name)
VALUES (?,?,?)
`

type CreateUserParams struct {
	CreateAt time.Time
	UpdateAt time.Time
	Name     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, arg.CreateAt, arg.UpdateAt, arg.Name)
}
