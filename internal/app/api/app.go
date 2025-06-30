package api

import (
	"orders/internal/config"
	"orders/internal/domain/storage/postgres"
	"orders/internal/router"
	"orders/internal/router/handlers"

	"braces.dev/errtrace"
)

func Run() error {
	cfg := config.InitConfig()

	storage, err := postgres.NewPostgres(cfg.Postgres)

	if err != nil {
		return errtrace.Wrap(err)
	}

	handlers := handlers.New(storage)

	if err := router.Run(cfg.ServerApi.Port, handlers); err != nil{
		return errtrace.Wrap(err)
	}
	return nil
}
