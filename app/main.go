package main

import (
	"fmt"
	"go-be-book/config"
	"go-be-book/models"
	"go-be-book/server"
)

func main() {
	db := config.DBConnect()
	models.DBMigrate(db)
	fmt.Println("Server started!")
	server.Run(db)
}
