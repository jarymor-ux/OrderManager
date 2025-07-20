package router

import (
	"orders/internal/router/handlers"

	"github.com/labstack/echo/v4"
)


func Run(port string, handlers handlers.Handlers) error {
	api := echo.New()

	v1 := api.Group("/api/v1")
	{
		v1.GET("/ping", handlers.Ping)
		v1.POST("/users", handlers.CreateUser)
	}

	if err := api.Start(port); err != nil {
		return err
	}
	return nil
}
