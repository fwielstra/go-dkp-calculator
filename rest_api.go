package main

import (
	"encoding/json"
	"net/http"
)

type getHandler struct {
	Name string
}

func (th *getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(r)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/get", &getHandler{Name: "hoi"})

	http.ListenAndServe(":1337", mux)
}
