package main

import (
	"testing"
)

func TestShouldAddUserToDkpService(t *testing.T) {
	state := state{users: make(map[string]User)}
	me := User{name: "silvan", points: 10}
	state.AddUser(me)
	meAgain := state.GetUser(me.name)
	if meAgain != me {
		t.Errorf("Error, expected %+v but was %+v", me, meAgain)
	}
}
