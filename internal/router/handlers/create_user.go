package handlers

import (
	"orders/internal/domain/entity"

	"braces.dev/errtrace"
	"github.com/labstack/echo/v4"
)

type RegisterUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handlers) CreateUser(e echo.Context) error {
	req := new(RegisterUserReq)
	
	if err := e.Bind(req); err != nil{
		return e.JSON(400,errtrace.Wrap(err))
	}

	user := new(entity.User).
		WithUsername(req.Username).
		WithPasswordHash(req.Password)

	if err := h.CreateUserUc.Run(*user); err != nil{
		return e.JSON(500, errtrace.Wrap(err))
	}

	return nil
}
