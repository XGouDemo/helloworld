package main

import (
	"fmt"
)

type Human struct {
	Job        string
	Aspiration string
	Body       string
	Love       string
}

func main6() {

	jerry := Human{
		Job:        "SAP Employee",
		Aspiration: "Travel Blogs",
		Body:       "Building",
		Love:       "Family",
	}

	tom := Human{"CEO of a DAX company", "Body Building", "Maintaining", "Family"}

	fmt.Println(jerry)
	fmt.Println(tom)
	transpersonate(&jerry) //pass the address
	fmt.Println(jerry)
}

func transpersonate(human *Human) {
	*human = Human{"Donald", "King", "perfect", "showing off"} //dereference the pointer
}
