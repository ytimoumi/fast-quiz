package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ytimoumi/fast-quiz/server/models"
	"io"
	"log"
	"net/http"
	"strconv"
)

// HandleQuestions encodes questions into JSON format and writes to the response.
func HandleQuestions(response http.ResponseWriter, request *http.Request) {
	log.Println("HandleQuestions =======> New request :", request.URL.Path)

	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		log.Println("Responding with HTTP status 400 Bad Request")
		response.WriteHeader(http.StatusBadRequest)
		response.Header().Set("Content-Type", "plain/text")
		io.WriteString(response, "Invalid ID."+"\n")
		return
	}

	Questions, questionsExist := models.QuestionsMap[id]

	if !questionsExist {
		log.Println("Responding with HTTP status 404 Not Found")
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// cache
	response.Header().Set("Cache-Control", "max-age=300")
	json.NewEncoder(response).Encode(Questions)

	log.Println("Responded with HTTP status 200 OK")

}
