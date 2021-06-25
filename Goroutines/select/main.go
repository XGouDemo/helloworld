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
		select {
		case c <- randomNum:
			fmt.Println("sent to channel")
		default:
			fmt.Println("dropped")
		}
		//sending faster than what the processing
		/*
			worker 1 got 59
			sent to channel
			worker 0 got 33
			sent to channel
			worker 2 got 43
			sent to channel
			worker 3 got 91
			sent to channel
			worker 4 got 2
			dropped
			dropped
			dropped
			dropped
			dropped
			dropped
			dropped
			dropped
			dropped
			dropped
			sent to channel
			worker 1 got 25
			sent to channel
			worker 0 got 151
		*/
		time.Sleep(time.Millisecond * 100)
	}
}
