package main


type User struct {
    name string
    points int
}

type Role struct {
  name string
  value int
}

type Event struct {
  name string
  roles []Role
  users []User
}
