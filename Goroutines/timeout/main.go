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
		case <-time.After(time.Millisecond * 10):
			fmt.Println("dropped by time out.")

			//if adding the following section, there will be no time out.
			/*remember that default fires immediately if no channel is
			available.
					default:
						fmt.Println("dropped by default.")*/
		}

		time.Sleep(time.Millisecond * 100)
	}
}
