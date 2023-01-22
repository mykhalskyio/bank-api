package main

import (
	"bank-api/config"
	"bank-api/internal/controller/http"
	service2 "bank-api/internal/service"
	"bank-api/internal/storage/postgres"
	"log"
)

// @title bank api
// @version 1.0
// @description REST API

// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.GetConfig()
	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	service := service2.NewService(db)
	router := http.NewRouter(service)
	log.Fatalln(router.Run(":8080"))
}
