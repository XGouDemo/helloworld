package main

import "fmt"

//definition of the function interface
type Add func(a int, b int) int

func process(addInt Add) int {
	//consumer of the abstract function Add
	return addInt(4, 3)
}

func main() {
	//implementation of the adding logic
	var processResult = process(func(a int, b int) int {
		return a*4 + b*5
	})

	fmt.Println(processResult)

	//a different implementation of the adding logic
	processResult = process(func(a int, b int) int {
		return a*7 + b*9
	})

	fmt.Println(processResult)

}
