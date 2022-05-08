package server

import (
	"github.com/gorilla/mux"
	"github.com/ytimoumi/fast-quiz/server/services"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/v1").Subrouter()
	// GET questions by id
	subRouter.HandleFunc("/questions/{id:\\d+}", services.HandleQuestions).Methods("GET")
	// POST answers
	subRouter.HandleFunc("/questions/{id:\\d+}/answers", services.HandleAnswers).Methods("POST")

	return router
}
