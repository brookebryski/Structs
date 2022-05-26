package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	Brooke := person{
		firstName: "Brooke",
		lastName:  "Bryski",
	}

	Alex := person{
		firstName: "Alex",
		lastName:  "Smith",
	}

	fmt.Println(Brooke)
	fmt.Println(Alex.firstName)

}
