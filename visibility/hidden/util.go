package secretUtil

import "fmt"

func print(str string) {
	fmt.Println("I can print for my client in this package.")
	fmt.Println("Nicely printing: '" + str + "'")
}

func Leak(str string) {
	print(str)
}
