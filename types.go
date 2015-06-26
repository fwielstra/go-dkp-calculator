package main

type user struct {
	Name   string
	Points int
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
