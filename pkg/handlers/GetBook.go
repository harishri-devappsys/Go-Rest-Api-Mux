package handlers

import (
	"encoding/json"
	"fmt"
	"gorestapi/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// for _, book := range mocks.Books {
	// 	if book.Id == id {
	// 		w.Header().Add("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)

	// 		json.NewEncoder(w).Encode(book)
	// 		break
	// 	}
	// }
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
