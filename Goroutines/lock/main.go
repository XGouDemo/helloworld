package main

import (
	"fmt"
	"sync"
	"time"
)

//not protected against writing by other goroutines
var counter = 0

//use the mutext to make sure the statements below lock() is executed in sequence by all goroutines
var lock sync.Mutex

/*
PS C:\Users\D058009\go\src\helloworld\Goroutines\lock> go run .\main.go
1
2
3
4
5
6
7
8
9
10
PS C:\Users\D058009\go\src\helloworld\Goroutines\lock>
*/

func main() {
	for i := 0; i < 10; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 1000)
}
func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}
