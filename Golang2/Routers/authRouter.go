package router

import (
	"github.com/gorilla/mux"
	"github.com/devesh/mongoapi/Controllers"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/users/signup", controller.SignupHandler).Methods("POST")
	router.HandleFunc("/users/login", controller.LoginHandler).Methods("POST")
}
