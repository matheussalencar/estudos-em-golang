package main

import "log"

func main() {
	var myString string
	myString = "green"

	log.Println("my string is set to", myString)
	changeUsingPointer(&myString)
	log.Println("after func call my string is set to", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"
	*s = newValue
}
