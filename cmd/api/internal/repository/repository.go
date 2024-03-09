package repository

import (
	"backend/cmd/api/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	OneMovie(id int) (*models.Movie, error)
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)

	GetUserByEmail(string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
