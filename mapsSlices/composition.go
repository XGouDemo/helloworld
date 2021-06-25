package main

import "fmt"

type Humana struct {
	Name       string
	Body       string
	Love       *Humana
	Job        string
	Aspiration string
}

func main5() {
	var tom Humana

	jerry := &Humana{"Jerry", "building", &tom, "developer", "reading"}
	tom = Humana{"Tom", "building", jerry, "developer", "reading"}

	fmt.Println(jerry)
	fmt.Println(tom)

}
