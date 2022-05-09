package services

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ytimoumi/fast-quiz/server/models"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

// HandleAnswers saves answers.
func HandleAnswers(response http.ResponseWriter, request *http.Request) {

	var answers []struct {
		AnswerID int `json:"answerId"`
	}

	log.Println("HandleAnswers =====> get new request", request.URL.Path)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		// if the ID is not an integer
		log.Println("Responding with HTTP status 400 Bad Request")
		response.WriteHeader(http.StatusBadRequest)
		response.Header().Set("Content-Type", "plain/text")
		io.WriteString(response, "ID must be an integer"+"\n")
		return
	}

	Questions, questionsExist := models.QuestionsMap[id]
	if !questionsExist {
		// invalid ID was supplied
		log.Println("Responding with HTTP status 404 Not Found")
		response.WriteHeader(http.StatusNotFound)
		return
	}

	// parsing JSON
	err = json.NewDecoder(request.Body).Decode(&answers)
	if err != nil {
		log.Println("Responding with HTTP status 400 Bad Request")
		response.WriteHeader(http.StatusBadRequest)
		response.Header().Set("Content-Type", "plain/text")
		return
	}

	// checking that all questions are answered
	totalQuestions := len(Questions)

	if totalQuestions != len(answers) {

		log.Println("Responding with HTTP status 400 Bad Request")
		response.WriteHeader(http.StatusBadRequest)
		response.Header().Set("Content-Type", "plain/text")
		io.WriteString(response, "Invalid amount of answers."+"\n")

		return
	}

	// calculate correct answers
	correctAnswers := 0
	for questionIx, answer := range answers {
		if answer.AnswerID == Questions[questionIx].CorrectAnswerID {
			correctAnswers++
		}
	}

	// respond with user's result
	Others := getOtherQuizzers(correctAnswers, totalQuestions)

	// update general statistics
	models.UsersAnsweredCorrectly[correctAnswers]++
	models.AnsweredUsers++

	res := struct {
		CorrectAnswers    int    `json:"correctAnswers"`
		ComparingToOthers string `json:"comparingToOthers"`
	}{
		CorrectAnswers:    correctAnswers,
		ComparingToOthers: "You were better then " + strconv.Itoa(Others) + "% of all quizzers",
	}
	json.NewEncoder(response).Encode(res)

	log.Println("Responded with HTTP status 200 OK")
}

// Returns by how many percentage current user has answered better than the rest
func getOtherQuizzers(correctAnswers int, totalQuestions int) int {
	if correctAnswers == 0 {
		// no correct answers, we are the worst :(*
		return 0
	}

	if correctAnswers == totalQuestions || models.AnsweredUsers == 0 {
		// answered all questions or nobody yet answered and we have
		// at least 1 correct answer, we are the best
		return 100
	}

	// calculating percentage
	usersAnsweredWorse := 0
	for i := 0; i <= correctAnswers; i++ {
		usersAnsweredWorse += models.UsersAnsweredCorrectly[i]
	}

	if usersAnsweredWorse == 0 {
		// nobody is worse than us, we are the worst again :(
		return 0
	}
	res := math.Ceil(float64(usersAnsweredWorse) / float64(models.AnsweredUsers) * 100.0)

	return int(res)
}
