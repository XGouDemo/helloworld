package main

import (
	//package name "relative file path"
	secretUtil "helloworld/visibility/hidden"
	util "helloworld/visibility/visible"
)

func main() {
	util.Print("go!")
	secretUtil.Leak("back!")

	//print is not exported by package setretUtil
	//secretUtil.print("steal!")
}
