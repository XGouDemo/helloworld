package main

import (
	"helloworld/interface/impl"
)

func main() {

	german := &impl.German{}
	chinese := &impl.Chinese{}

	german.Run(5)
	chinese.Run(5)
}
