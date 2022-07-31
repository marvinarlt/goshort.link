package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewFrontend(router *mux.Router) {
	router.HandleFunc("/", serveIndex)
	router.HandleFunc("/legal-notice", serveIndex)
	router.HandleFunc("/privacy-policy", serveIndex)
}

func serveIndex(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./frontend/dist/index.html")
}
