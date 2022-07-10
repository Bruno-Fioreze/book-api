package main

import (
	"github.com/bruno-fioreze/book_api/server"
	"github.com/hyperyuri/webapi-with-go/database"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
