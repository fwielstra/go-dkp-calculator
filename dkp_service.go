package main

type state struct {
	users  map[string]User
	events []Event
}

func (s *state) AddUser(user User) {
	s.users[user.name] = user
}

func (s *state) GetUser(name string) User {
	return s.users[name]
}
