package main

import (
	"fmt"
)

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)

	lookup["lihong"] = 9987
	power, exists = lookup["lihong"]
	// prints 9987 true
	fmt.Println(power, exists)

	// returns 2
	total := len(lookup)
	fmt.Println(total)
	// has no return, can be called on a nonexisting key
	delete(lookup, "goku")
	// returns 1
	total = len(lookup)
	fmt.Println(total)
}
