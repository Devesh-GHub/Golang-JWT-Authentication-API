package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

func JWTTestRoute(router *mux.Router) {
	router.HandleFunc("/jwt-test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("JWT Route Working"))
	}).Methods("GET")
}