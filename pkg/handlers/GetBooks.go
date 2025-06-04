package handlers

import (
	"encoding/json"
	"fmt"
	"gorestapi/pkg/models"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(mocks.Books)
	var books []models.Book

    if result := h.DB.Find(&books); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(books)
}