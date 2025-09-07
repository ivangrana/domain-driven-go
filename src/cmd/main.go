package main

import (
	"log"
	"net/http"
	"os"

	"domain-driven-go/src/internal/application"
	"domain-driven-go/src/internal/domain/service"
	"domain-driven-go/src/internal/infrastructure/persistence"
	"domain-driven-go/src/presentation/api"
)

func main() {
	// Get MongoDB connection details from environment variables
	mongodbURI := os.Getenv("MONGODB_URI")
	mongodbDB := os.Getenv("MONGODB_DB")

	// Create dependencies
	userRepo, err := persistence.NewMongoDBUserRepository(mongodbURI, mongodbDB, "users")
	if err != nil {
		log.Fatal(err)
	}
	userService := service.NewUserService(userRepo)
	userApp := application.NewUserApplication(userService)
	userHandler := api.NewUserHandler(userApp)
	routes := api.RegisterRoute(userHandler)

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
