package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `

[
	{ 
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color":"black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color":"black",
		"has_dog": false
	}
]`
	var unmasrhalled []Person

	err := json.Unmarshal([]byte(myJson), &unmasrhalled)

	if err != nil {
		log.Println("Error unmarshlling json", err)
	}

	log.Printf("unmarshalled: %v", unmasrhalled)

	// write json from struct

	var mySlice []Person

	var m1 Person
	m1.FirstName = "wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	m1.FirstName = "matheus"
	m1.LastName = "santos"
	m1.HairColor = "yellow"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	log.Println(mySlice)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")

	if err != nil {
		log.Println("Error marshalling: ", err)
	}

	fmt.Println(string(newJson))

}
