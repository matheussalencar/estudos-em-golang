package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Maps
	myMap := make(map[string]User)

	me := User{
		FirstName: "Matheus",
		LastName:  "Santos",
	}

	myMap["me"] = me

	log.Println(me.FirstName, me.LastName)

	// slices
	var mySlice []string
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)

	log.Println(numbers[0:3])

	mySlice = append(mySlice, "Matheus")
	mySlice = append(mySlice, "Santos")
	mySlice = append(mySlice, "Alencar")

	sort.Strings(mySlice)

	log.Println(mySlice)
}
