package main

import (
	"fmt"
	"log"
)

func main() {
	isAdmin := "ADMIN"
	role := "TEST"

	if adminTasks(isAdmin) {
		fmt.Println("You have permission to do that")
	} else {
		fmt.Println("else statement")
	}

	switch role {
	case "USER":
		log.Println("USER SYSTEM")

	case "ADMIN":
		log.Println("ADMIN SYSTEM")

	default:
		log.Println("nothing to do")
	}
}

func adminTasks(role string) bool {
	if role == "ADMIN" {
		fmt.Println("Welcome to admin system")
		return true
	} else {
		fmt.Println("You dont hava permission to access admin system")
		return false
	}
}
