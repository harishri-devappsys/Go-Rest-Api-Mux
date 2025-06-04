package handlers

import (
	"encoding/json"
	"fmt"

	// "gorestapi/pkg/mocks"
	"gorestapi/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock Books
	// for index, book := range mocks.Books {
	//     if book.Id == id {
	//         mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

	//         w.Header().Add("Content-Type", "application/json")
	//         w.WriteHeader(http.StatusOK)
	//         json.NewEncoder(w).Encode("Deleted")
	//         break
	//     }
	// }

	//Postgre implementation
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	h.DB.Delete(&book)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
