package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthday    time.Time
}

func main() {
	user := User{
		FirstName:   "Matheus",
		LastName:    "Santos",
		PhoneNumber: "(32)-XXX-XXX",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber)
}
