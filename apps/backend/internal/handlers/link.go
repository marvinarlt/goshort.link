package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/marvinarlt/goshort.link/internal/models"
	"github.com/marvinarlt/goshort.link/internal/services"
)

// Create new id and save it with the URL in the database
func NewLink(response http.ResponseWriter, request *http.Request) {
	var link models.Link

	err := json.NewDecoder(request.Body).Decode(&link)

	request.Body.Close()

	if nil != err {
		http.Error(response, "Invalid JSON", http.StatusBadRequest)
		return
	}

	link.Id = services.UniqueId(12)

	services.DB.Select("id", "url").Create(&link)

	corsHost := os.Getenv("BACKEND_CORS_ORIGIN")

	// if "production" == os.Getenv("ENVIRONMENT") {
	// 	corsHost = os.Getenv("CORS_HOST_PROD")
	// }

	response.Header().Set("Access-Control-Allow-Origin", corsHost)
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("Access-Control-Allow-Methods", "POST")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(link)
}

// Redirect based on the id
func GetLink(response http.ResponseWriter, request *http.Request) {
	request.Body.Close()

	link := models.Link{
		Id: mux.Vars(request)["id"],
	}

	result := services.DB.Take(&link)

	if nil != result.Error {
		http.Redirect(response, request, "/", http.StatusSeeOther)
		return
	}

	link.LatestUse = time.Now().UTC()
	services.DB.Save(&link)

	http.Redirect(response, request, link.Url, http.StatusSeeOther)
}
