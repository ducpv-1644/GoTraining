package handler

import (
	"encoding/json"
	"fmt"
	"go-be-book/books"
	"go-be-book/books/repository"
	bookUsecase "go-be-book/books/usecase"
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

func newUsecase() books.Usecase {
	return bookUsecase.NewBooksUsecase(
		repository.NewBookRepository(),
	)
}

func (a *BookHandler) ListBook(w http.ResponseWriter, r *http.Request) {
	db := config.DBConnect()
	books, err := newUsecase().ListBooks(db)
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
	book_repository.DeleteBook(db, bookId)

	type response struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}
	resp := response{}
	resp.Code = 200
	respondWithJSON(w, http.StatusOK, resp)
}

func (a *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	db := config.DBConnect()
	book_repository := repository.NewBookRepository()
	type response struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	}
	resp := response{}

	vars := mux.Vars(r)
	bookId, error := vars["id"]
	if !error {
		resp.Code = 400
		resp.Message = "id is missing in parameters!"
		respondWithJSON(w, http.StatusOK, resp)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.Code = 400
		resp.Message = "Get body info failed!"
		respondWithJSON(w, http.StatusOK, resp)
		return
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	book := book_repository.UpdateBook(db, bookId, keyVal["title"], keyVal["author"])
	respondWithJSON(w, http.StatusOK, book)
}
