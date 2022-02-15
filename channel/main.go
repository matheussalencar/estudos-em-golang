package main

import (
	"github.com/matheussalencar/estudos-em-golang/helpers"
	"log"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}
func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	myNumber := <-intChan

	log.Println(myNumber)
}
