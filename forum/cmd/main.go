package main

import (
	"forum/internal/handler"
	"forum/internal/service"
	"forum/internal/storage"
	"log"
	"net/http"
)

func main() {
	database := storage.InitDB()
	storages := storage.NewStorage(database)
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)
	handlers.InitRoutes()
	log.Println("Running a web server on http://localhost:8080")
	http.ListenAndServe(":8080", handlers.Mux)
}
