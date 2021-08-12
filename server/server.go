package server

import (
	"go-be-book/models"
	"go-be-book/server/handler"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)



func Run(db *gorm.DB) {
	book_handler := handler.BookHandler{}
	router := mux.NewRouter()
	models.DBMigrate(db)

	router.HandleFunc("/", book_handler.ListBook).Methods("GET")
	router.HandleFunc("/books", book_handler.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", book_handler.RetrieveBook).Methods("GET")
	router.HandleFunc("/book/{id}/delete", book_handler.DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
