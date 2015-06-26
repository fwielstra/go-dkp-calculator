package main

import (
  "encoding/json"
  "io"
)

type user struct {
	Name   string
	Points int
}

func ParseUser(r io.Reader) user {
  var user user
  json.NewDecoder(r).Decode(&user)
  return user
}

type role struct {
	name  string
	value int
}

type event struct {
	name  string
	roles []role
	users []user
}
