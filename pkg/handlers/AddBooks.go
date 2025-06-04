package handlers

import (
	"encoding/json"
	"fmt"
	"gorestapi/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// book.Id = rand.Intn(100)
	// mocks.Books = append(mocks.Books, book)
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
