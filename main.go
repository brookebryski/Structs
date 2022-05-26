package main

import "fmt"

type person struct {
	fistName string
	lastName string
}

func main() {
	Brooke := person{
		fistName: "Brooke",
		lastName: "Bryski",
	}

	fmt.Println(Brooke)
}