package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var appState state

type getHandler struct {
}

func (th *getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user{Name: "henk", Points: 1337})
}

type getRequestHandler func() interface{}

func handle(path string, fn getRequestHandler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(fn())
	})
}

func main() {
	appState = state{users: make(map[string]user)}
	appState.AddUser(user{Name: "Henk", Points: 1337})

	handle("/user", func() interface{} {
		return appState.GetUsers()
	})

	log.Fatal(http.ListenAndServe(":1337", nil))
}
