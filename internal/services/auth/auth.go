package auth

import (
	"github.com/labstack/echo/v4"
	"orders/internal/domain/models"
)

// TODO: |rexiizet| - implement auth service
type Auth struct {
	iUser IUser
}

type IUser interface {
	RegisterUser(
		ctx echo.Context,
		user models.User,
	) (id int64, err error)
	AuthUser(
		ctx echo.Context,
		login string,
	) (models.User, error)
}

func New(
	iUser IUser,
) (*Auth, error) {

	return &Auth{iUser: iUser}, nil
}
