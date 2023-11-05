package main

import (
	"time"

	"github.com/Shin198/go-be-rss/internal/database"
)

type User struct {
	ID       int32     `json:"id"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	Name     string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:       dbUser.ID,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
		Name:     dbUser.Name,
	}
}
