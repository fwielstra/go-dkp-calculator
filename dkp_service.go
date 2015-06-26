package main

type state struct {
	users  map[string]user
	events []event
}

func (s *state) AddUser(user user) {
	s.users[user.Name] = user
}

func (s *state) GetUser(name string) user {
	return s.users[name]
}

func (s *state) GetUsers() []user {
	result := make([]user, 0, len(s.users))

	for  _, user := range s.users {
		result = append(result, user)
	}

	return result
}
