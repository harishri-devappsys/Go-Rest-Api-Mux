package handlers

import (
	"encoding/json"
	"fmt"
	"gorestapi/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	// var updatedBook models.Book
	// json.Unmarshal(body, &updatedBook)

	// for index, book := range mocks.Books {
	// 	if book.Id == id {
	// 		book.Title = updatedBook.Title
	// 		book.Author = updatedBook.Author
	// 		book.Desc = updatedBook.Desc

	// 		mocks.Books[index] = book
	// 		w.Header().Add("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)

	// 		json.NewEncoder(w).Encode("Updated")
	// 		break
	// 	}
	// }
	var updatedBook models.Book
	json.Unmarshal(body, &updatedBook)

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Desc = updatedBook.Desc

	h.DB.Save(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
