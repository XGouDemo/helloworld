package main1

import "fmt"

type Human struct {
	Name       string
	Body       string
	Love       *Human
	Job        string
	Aspiration string
}

func main() {
	var tom Human

	jerry := &Human{"Jerry", "building", &tom, "developer", "reading"}
	tom = Human{"Tom", "building", jerry, "developer", "reading"}

	fmt.Println(jerry)
	fmt.Println(tom)

}
