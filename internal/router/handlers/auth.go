package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"orders/internal/domain/requests/auth"
	authService "orders/internal/services/auth"
)

func SignIn(ctx echo.Context) error {
	var input auth.LoginRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid request format",
		})
	}

	//TODO: |rexiizet| - сделать authClient и разместить его в internal/ (clients or services)

	return ctx.JSON(http.StatusOK, map[string]any{
		"token":    user.Token,
		"userId":   user.UserId,
		"username": user.Username,
		"phone":    user.Phone,
		"email":    user.Email,
	})
}

func SignUp(ctx echo.Context) error {
	var input auth.RegisterRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": "invalid request format",
		})
	}

	//TODO: |rexiizet| - сделать authClient и разместить его в internal/ (clients or services)

	log.Println("Register method succeeded")

	return ctx.JSON(http.StatusOK, map[string]any{
		"token":    user.Token,
		"userId":   user.UserId,
		"username": user.Username,
		"phone":    user.Phone,
		"email":    user.Email,
	})
}
