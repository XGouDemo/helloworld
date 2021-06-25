package main3

import "fmt"

type Human struct {
	Name       string
	Body       string
	Love       string
	Job        string
	Aspiration string
}

func main() {
	jerry := &Human{"Jerry", "building", "wife", "developer", "reading"}
	tom := Human{"Tom", "building", "wife", "developer", "reading"}
	brad := new(Human)
	brad.Name = "Brad"

	//using the dot, both a variable and its pointer could refer to the instance of the structure.
	jerry.Body = "transforming" //effect is same as the line below
	tom.Body = "transforming"   //effect is same as the line above

	fmt.Println(*jerry)
	fmt.Println(tom)
	fmt.Println(brad)

	jerry.boost()

	fmt.Println(*jerry)
}

func (h *Human) boost() {
	fmt.Println(h)
	h.Body = "boosting"
}
