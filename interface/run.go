package main

import (
	"interface/impl"
)

func main() {

	german := &impl.German{}
	chinese := &impl.Chinese{}

	german.Run(5)
	chinese.Run(5)
}
