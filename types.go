package main

type user struct {
	name   string
	points int
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
