package router

import (
	"orders/internal/router/handlers"

	"github.com/labstack/echo/v4"
)

func Run(port string) error {
	api := echo.New()

	v1 := api.Group("/api/v1")
	{
		v1.GET("/ping", handlers.Ping)
	}

	auth := v1.Group("/auth")
	{
		auth.POST("/sign-up", handlers.SignUp)
		auth.POST("/sign-in", handlers.SignIn)
	}

	if err := api.Start(port); err != nil {
		return err
	}
	return nil
}
