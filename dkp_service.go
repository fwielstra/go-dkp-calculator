package main

type state struct {
	users  map[string]user
	events []event
}

func (s *state) AddUser(user user) {
	s.users[user.name] = user
}

func (s *state) GetUser(name string) user {
	return s.users[name]
}
