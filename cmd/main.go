package main

import (
	"gorestapi/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handlers.Getbook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)
	log.Println("Api is running")
	http.ListenAndServe(":4000", router)

	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
