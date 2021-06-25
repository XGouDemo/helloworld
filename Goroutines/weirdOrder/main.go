package main

import (
	"fmt"
	"time"
)

//not protected against writing by other goroutines
var counter = 0

/*
PS C:\Users\D058009\go\src\helloworld\Goroutines\weirdOrder> go run .\main.go
158
158
158
PS C:\Users\D058009\go\src\helloworld\Goroutines\weirdOrder>
*/

func main() {
	for i := 0; i < 10; i++ {
		go incr()
	}
	time.Sleep(time.Microsecond * 1)
}
func incr() {
	counter++
	go incr()
	time.Sleep(time.Microsecond * 100)

	//printing results depending on the time when the coutner is read. It is not deterministic from the static coding
	fmt.Println(counter)
}
