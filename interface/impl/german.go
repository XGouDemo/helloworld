package impl

import (
	"fmt"
)

type German struct{}

func (g German) Run(speed int) {
	fmt.Print("German at speed: ")
	fmt.Println(speed)
}
