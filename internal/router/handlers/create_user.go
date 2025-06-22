package handlers

import (
	"orders/internal/domain/responses"

	"github.com/labstack/echo/v4"
)

//TODO:IMPLEMENT!!
func CreateUser(e echo.Context) error {
	return e.JSON(500, responses.JsonResponse("msg", "IMPLEMENT ME"))
}
