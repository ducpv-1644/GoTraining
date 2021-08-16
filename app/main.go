package main

import (
	// "fmt"
	"go-be-book/config"
	"go-be-book/models"
	"go-be-book/server"
	"sync"

	// "gorm.io/gorm"
)

const num_workers = 1

func main() {
	var wg sync.WaitGroup
	wg.Add(num_workers)

	db := config.DBConnect()
	models.DBMigrate(db)
	go server.Run(&wg)

	wg.Wait()
}
