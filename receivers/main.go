package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstname() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "Matheus"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myvar is set to", myVar.printFirstname())
	log.Println("myvar2 is set to", myVar2.printFirstname())
}
