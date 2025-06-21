package api

import (
	"orders/internal/config"
	"orders/internal/router"
)

func Run() error {
	cfg := config.InitConfig()
	if err := router.Run(cfg.ServerApi.Port); err != nil{
		return  err
	}
	return nil
}
