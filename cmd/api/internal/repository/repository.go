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
	InsertMovie(movie models.Movie) (int, error)
	UpdateMovie(movie models.Movie) error
	DeleteMovie(id int) error

	AllGenres() ([]*models.Genre, error)
	UpdateMovieGenres(id int, genreIDs []int) error

	GetUserByEmail(string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
