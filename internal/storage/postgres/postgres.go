package postgres

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"orders/internal/domain/models"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) (*Storage, error) {
	const op = "postgres.NewStorage"

	return &Storage{db: db}, nil
}

func (s *Storage) RegisterUser(ctx echo.Context, user models.User) (int64, error) {
	const op = "postgres.NewStorage.SaveUser"

	//TODO: |rexiizet| - implement query to DB

	return 0, nil
}

func (s *Storage) AuthUser(ctx echo.Context, login string) (models.User, error) {
	const op = "postgres.NewStorage.GetUser"

	//TODO: |rexiizet| - implement query to DB

	return models.User{}, nil
}
