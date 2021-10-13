package main

import (
	// "hello/src/server"
	"hello/src/server"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := server.NewPlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":5000", server))
}
