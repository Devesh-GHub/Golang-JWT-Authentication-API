package router

import (
	"github.com/gorilla/mux"
	"github.com/devesh/mongoapi/Controllers"
	"github.com/devesh/mongoapi/Middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/register", controller.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/api/login", controller.LoginUserHandler).Methods("POST")

	r.Use(middleware.Authenticate)

	return r
}