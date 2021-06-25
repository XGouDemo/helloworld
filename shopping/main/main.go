package main

import (
	"fmt"
	"shopping"
)

//main.go:5:2: package shopping is not in GOROOT (C:\Program Files\Go\src\shopping)
//resolution: turnning off GO Module by:  go env -w GO111MODULE=off
func main() {
	fmt.Println(shopping.PriceCheck(4343))
}
