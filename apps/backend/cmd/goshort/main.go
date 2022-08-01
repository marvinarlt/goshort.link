package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marvinarlt/goshort.link/internal/consts"
	"github.com/marvinarlt/goshort.link/internal/routes"
	"github.com/marvinarlt/goshort.link/internal/services"
)

func main() {
	services.DbConnect()
	services.RemoveUnused()

	router := mux.NewRouter()
	routes.NewShortener(router)
	routes.NewFrontend(router)
	http.Handle("/", router)

	fmt.Println(consts.PREFIX, "Server started")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
