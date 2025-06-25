package handlers

import (
	"net/http"
	"orders/internal/domain/responses"

	"github.com/labstack/echo/v4"
)

func Ping(e echo.Context) error {
	return e.JSON(http.StatusOK, responses.JsonResponse("msg", "pong"))
}
