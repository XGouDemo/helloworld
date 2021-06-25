package main

import (
	"fmt"
)

func main() {
	scores := []int{1, 2, 3, 4, 5}
	slice := scores[2:4]
	slice[0] = 999
	fmt.Println(scores) //[1 2 999 4 5]

	slice[1] = 999
	fmt.Println(scores) //[1 2 999 999 5]

	//slice[2] = 999 --> panic: runtime error: index out of range [2] with length 2
	//The [X:Y] syntax creates a slice of scores, starting from index 2 up until (but not including) index 4.
	fmt.Println(len(slice)) //2

}
