package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var appState state

type requestHandler func(r *http.Request) interface{}

func handle(path string, fn requestHandler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(fn(r))
	})
}

func main() {
	appState = state{users: make(map[string]user)}

	handle("/user", func(r *http.Request) interface{} {
		switch r.Method {
		case "GET": return appState.GetUsers()
		case "POST":
			user := ParseUser(r.Body)
			appState.AddUser(user)
			return user
		}

		return nil
	})

	log.Fatal(http.ListenAndServe(":1337", nil))
}
