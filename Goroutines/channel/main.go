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
		//due to the sleeping period, the income will not be picked up without buffering in channel
		time.Sleep(time.Millisecond * 1500)
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
		//sending faster than what the processing
		time.Sleep(time.Millisecond * 50)
	}
}
