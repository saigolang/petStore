package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"petStore/backend"
	"petStore/logger"
)

func main() {

	// Initializing logging
	logger.NewLogger()
	router := mux.NewRouter()

	router.HandleFunc("/pets", backend.GetPets).Methods("GET")
	router.HandleFunc("/pet/{id}", backend.GetPetById).Methods("GET")
	router.HandleFunc("/create", backend.CreatePetResource).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Server started")

	//start and listen to requests
	log.Fatal(http.ListenAndServe(":8082", router))
}
