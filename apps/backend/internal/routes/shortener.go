package routes

import (
	"github.com/gorilla/mux"
	"github.com/marvinarlt/goshort.link/internal/handlers"
)

// Create new link routes
func NewShortener(router *mux.Router) {
	router.HandleFunc("/api/link", handlers.NewLink).Methods("POST")
	router.HandleFunc("/{id:[a-zA-Z0-9]{12}}", handlers.GetLink).Methods("GET")
}
