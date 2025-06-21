package main

import (
	"log"
	"orders/internal/app/api"
)

func main() {
	if err := api.Run(); err != nil{
		log.Panic(err)
	}
}
