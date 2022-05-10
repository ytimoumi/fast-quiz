package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Init() {
	router := NewRouter()
	log.Println("\n Started server on port http://localhost:9444")
	if err := http.ListenAndServe(":9444", router); err != nil {
		log.Println("Listening failed")
		fmt.Print(err)
		os.Exit(1)
	}
}
