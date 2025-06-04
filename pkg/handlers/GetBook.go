package handlers

import (
	"encoding/json"
	"gorestapi/pkg/mocks"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Getbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, book := range mocks.Books {
		if book.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
