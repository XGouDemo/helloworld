package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "the spice must flow"

	fmt.Println(haystack[5:])                     //pice must flow
	fmt.Println(strings.Index(haystack[5:], " ")) //4
	fmt.Println(strings.Index(haystack[5:], "p")) //0
	fmt.Println(strings.Index(haystack[5:], "m")) //5
}
