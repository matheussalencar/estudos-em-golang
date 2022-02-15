package main

import "fmt"

func main() {

	mySlice := []string{"Matheus", "Maria", "John", "javier"}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < len(mySlice); i++ {
		if mySlice[i] == "Maria" {
			fmt.Println("Achei a ", mySlice[i])
			i = i + 1
		}
		fmt.Println(mySlice[i])
	}

	for _, mySlice := range mySlice {
		fmt.Println(mySlice)
	}

}
