package handlers

import (
	"orders/internal/domain/responses"

	"github.com/labstack/echo/v4"
)


func Ping(e echo.Context) error{
	return e.JSON(200, responses.JsonResponse("msg", "pong"))
}
