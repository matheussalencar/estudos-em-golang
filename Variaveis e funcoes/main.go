package main

import "fmt"

var myName string // package level variable

func main() {
	fmt.Println("you are ready to go!")

	var whatToSay string
	var i int

	myName = "matheus"
	fmt.Println(myName)

	whatToSay = "I only know go!"

	fmt.Println(whatToSay)

	i = 10

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThing := saySomething()

	fmt.Println(whatWasSaid, theOtherThing)
}

func saySomething() (string, string) { // uma funcao pode retornar mais de um valor
	return "something", "else"
}
