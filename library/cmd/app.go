package main

import (
	"log"

	"github.com/CesarDelgadoM/microservices-go/library/config"
	"github.com/CesarDelgadoM/microservices-go/library/internal/adapters/repositories"
	"github.com/CesarDelgadoM/microservices-go/library/server"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfigPostgres("../config")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	db := repositories.ConnectPostgres(&config)

	engine := gin.Default()

	err = server.NewServer(db, engine).Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
