package impl

import (
	"fmt"
)

type Chinese struct{}

func (C Chinese) Run(speed int) {
	fmt.Print("Chinese at speed: ")
	fmt.Println(speed)
}
