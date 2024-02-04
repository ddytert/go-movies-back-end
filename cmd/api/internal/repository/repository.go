package repository

import (
	"backend/cmd/api/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
