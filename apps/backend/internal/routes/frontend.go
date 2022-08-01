package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewFrontend(router *mux.Router) {
	fileServer := http.FileServer(http.Dir("./dist"))
	router.PathPrefix("/").Handler(fileServer).Methods("GET")
}
