package main

import (
	"github.com/ronanzindev/book-api-golang/database"
	"github.com/ronanzindev/book-api-golang/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
