package main

import (
	"net/http"
	"fmt"
	"go-be-book/config"
	"go-be-book/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println("Running!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dbconfig, err := config.NewDBSetting()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Not create DCSetting!")
			fmt.Fprintf(w, "Not create DCSetting!")
			return
		}
		dsn := "host=" + dbconfig.DBHost + " port=5432 user=postgres dbname=go_training password=postgres sslmode=disable"
		fmt.Println(dsn)
		db, err := gorm.Open("postgres", dsn)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Db connect failed!")
			fmt.Fprintf(w, "Db connect failed!")
			return
		}
		db.AutoMigrate(&models.Book{})
		defer db.Close()

		fmt.Fprintf(w, "Db Connected!")
	})

	http.ListenAndServe(":8000", nil)
}
