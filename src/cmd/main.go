package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"domain-driven-go/src/internal/application"
	"domain-driven-go/src/internal/domain/service"
	"domain-driven-go/src/internal/infrastructure/persistence"
	"domain-driven-go/src/presentation/api"
)

func main() {
	// Create dependencies
	userRepo := persistence.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userApp := application.NewUserApplication(userService)
	userHandler := api.NewUserHandler(userApp)

	// Create router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
