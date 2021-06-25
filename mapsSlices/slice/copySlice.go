package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main1() {
	scores := make([]int, 10)
	for i := 0; i < 10; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	sort.Ints(scores) //[59 81 81 300 318 425 456 540 847 887]
	fmt.Println(scores)

	smallest := make([]int, 5)
	copy(smallest, scores[:5])
	fmt.Println(smallest) //[59 81 81 300 318]

	smallest = make([]int, 5)
	copy(smallest[0:4], scores[:5])
	fmt.Println(smallest) //[59 81 81 300 0]

	smallest = make([]int, 5)
	copy(smallest[1:4], scores[:5])
	fmt.Println(smallest) //[0 59 81 81 0]

	smallest = make([]int, 6)
	copy(smallest[1:4], scores[:5])
	fmt.Println(smallest) //[0 59 81 81 0 0]

	smallest = make([]int, 6)
	copy(smallest[1:4], scores[:7])
	fmt.Println(smallest) //[0 59 81 81 0 0]
}
