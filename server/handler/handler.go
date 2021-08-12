package handler

import (
	"encoding/json"
	"fmt"
	"go-be-book/books/repository"
	"go-be-book/config"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type BookHandler struct{}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (a *BookHandler) ListBook(w http.ResponseWriter, r *http.Request) {
	db := config.DBConnect()
	bookRepository := repository.NewBookRepository()
	books, err := bookRepository.ListBook(db)
	if err != nil {
		fmt.Println("Get list Book failed !")
	}

	respondWithJSON(w, http.StatusOK, books)
}

func (a *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Create Book failed !")
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	db := config.DBConnect()
	book_repository := repository.NewBookRepository()
	books, err := book_repository.CreateBook(db, keyVal["title"], keyVal["author"])
	if err != nil {
		fmt.Println("Create Book failed !")
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func (a *BookHandler) RetrieveBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, error := vars["id"]
	if !error {
		fmt.Println("id is missing in parameters")
		return
	}

	db := config.DBConnect()
	book_repository := repository.NewBookRepository()
	book, err := book_repository.RetrieveBook(db, bookId)
	if err != nil {
		fmt.Println("Create Book failed !")
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

func (a *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, error := vars["id"]
	if !error {
		fmt.Println("id is missing in parameters")
		return
	}

	db := config.DBConnect()
	book_repository := repository.NewBookRepository()
	id, err := book_repository.DeleteBook(db, bookId)
	if err != nil {
		fmt.Println("Create Book failed !")
		fmt.Println(id)
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}
