package routes

import (
	"github.com/gorilla/mux"
	"github.com/marvinarlt/goshort/internal/handlers"
)

// Create new link routes
func NewShortener(router *mux.Router) {
	router.HandleFunc("/api/link", handlers.NewLink).Methods("POST")
	router.HandleFunc("/{id}", handlers.GetLink).Methods("GET")
}
