package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (w Worker) process(c chan int32) {
	//It waits until data is available then “processes” it.
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}

func main() {
	c := make(chan int32)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		randomNum := rand.Int31n(200)
		c <- randomNum
		time.Sleep(time.Millisecond * 500)
	}
}
