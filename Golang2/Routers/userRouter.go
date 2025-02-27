package router

import (
	"github.com/gorilla/mux"
	"github.com/devesh/mongoapi/Controllers"
	"github.com/devesh/mongoapi/Middleware"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/users/{user_id}", controller.GetUserHandler).Methods("GET")
	router.Use(middleware.Authenticate)
}