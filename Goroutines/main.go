package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   1
	   processing
	   2
	   3
	   processing
	   processing
	   4
	   done
	*/
	go process()
	fmt.Println("1")
	time.Sleep(time.Microsecond * 1) // this is bad, don't do this!
	fmt.Println("2")
	time.Sleep(time.Microsecond * 1) // this is bad, don't do this!
	fmt.Println("3")
	time.Sleep(time.Microsecond * 1) // this is bad, don't do this!
	fmt.Println("4")
	time.Sleep(time.Microsecond * 1) // this is bad, don't do this!
	fmt.Println("done")
}
func process() {
	for {
		fmt.Println("processing")
		time.Sleep(time.Microsecond * 1) // this is bad, don't do this!
	}
}
