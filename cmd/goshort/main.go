package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/marvinarlt/goshort/internal/consts"
	"github.com/marvinarlt/goshort/internal/routes"
	"github.com/marvinarlt/goshort/internal/services"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	envErr := godotenv.Load()

	if nil != envErr {
		log.Fatalln(consts.PREFIX, "Could not load .env file")
	}

	services.DbConnect()
	services.RemoveUnused()

	router := mux.NewRouter()
	routes.NewShortener(router)
	http.Handle("/", router)

	fmt.Println(consts.PREFIX, "Server started")
	fmt.Println(consts.PREFIX, "Listening on port 8080")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
