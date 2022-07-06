package main

import "github.com/bruno-fioreze/book_api/server"

func main() {
	server := server.NewServer()
	server.Run()
}
