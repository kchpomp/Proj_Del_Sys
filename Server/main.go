package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	// Формат: "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName"
	server := newServer("postgres://postgres:d19995678@localhost:5432/Delivery_System")
	server.createRouter()

	ln, err := net.Listen("tcp", ":8181")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(ln, server.router))
}
