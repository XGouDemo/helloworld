package main

import (
	types "error/declaration"
	"fmt"
	"io"
	"strconv"
)

func main() {
	for {
		fmt.Println(scanner())
	}
}

func scanner() error {
	var input int
	fmt.Print("Input ->")
	argNum, err := fmt.Scan(&input, &input, &input)
	if err == io.EOF {
		fmt.Println("no more input!")
	}

	fmt.Println("Your input count was: " + strconv.Itoa(argNum))
	if argNum > 1 {
		err = types.MinorError
	} else {
		err = types.MajorError
	}

	return err
}
