package storage

import (
	"database/sql"
	"skillbox/middleware"
	"skillbox/pkg/models"
)

type Storage interface {
	Get(key int) *models.User
	Save(u *models.User) (int, error)
	AllUsers() []models.User
	Delete(id int) error
	Update(id int, u *models.User) (*models.User, error)
	DeleteFriend(id int)
}

func NewStorage() *sql.DB {
	db, err := sql.Open("sqlite3", "./storage/persons.db")
	middleware.CheckErr(err)
	return db
}
