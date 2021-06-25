package main

import (
	"fmt"
)

func main() {
	/*
		a is an integer
		b is an integer
		13
	*/
	result := add(5, 8)
	fmt.Println(result)
	//a is an integer
	//ERROR: Type Mismatch
	result = add(5, 8.3)
	fmt.Println(result)
}

func add(a interface{}, b interface{}) interface{} {
	var c interface{}

	switch a.(type) {
	case int:
		fmt.Println("a is an integer")
		switch b.(type) {
		case int:
			fmt.Println("b is an integer")
			c = a.(int) + b.(int)
		default:
			return "ERROR: Type Mismatch"
		}
	default:
		return "ERROR: Type Mismatch"
	}
	return c.(int)
}
