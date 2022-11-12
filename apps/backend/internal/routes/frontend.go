package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewFrontend(router *mux.Router) {
	fileServer := http.FileServer(http.Dir("./public"))
	router.PathPrefix("/").Handler(fileServer).Methods("GET")
}
