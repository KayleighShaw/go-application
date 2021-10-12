package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// func PlayerServer(w http.ResponseWriter, r *http.Request) {
// 	player := strings.TrimPrefix(r.URL.Path, "/players/") //returns the path of the request (r.URL.Path), strings.TrimPrefix: trims away "/players/" to get the requested player

// 	fmt.Fprint(w, GetPlayerScore(player))
// }

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
