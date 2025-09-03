package api

import (

	"github.com/gorilla/mux"

)

func RegisterRoute(userHandler *UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	return r
}
