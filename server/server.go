package server

import (
	"fmt"
	"go-be-book/server/handler"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)



func Run(wg *sync.WaitGroup) {
        book_handler := handler.BookHandler{}
        router := mux.NewRouter()

        defer wg.Done()

        router.HandleFunc("/", book_handler.ListBook).Methods("GET")
        router.HandleFunc("/books", book_handler.CreateBook).Methods("POST")
        router.HandleFunc("/book/{id}", book_handler.RetrieveBook).Methods("GET")
        router.HandleFunc("/book/{id}/delete", book_handler.DeleteBook).Methods("DELETE")
        router.HandleFunc("/book/{id}/update", book_handler.UpdateBook).Methods("PUT")
        router.HandleFunc("/books/create-combo", book_handler.CreateComboBook).Methods("POST")

        fmt.Println("Server started port 8000!")
        http.ListenAndServe(":8000", router)
}
