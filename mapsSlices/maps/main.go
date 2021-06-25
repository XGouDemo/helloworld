package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Romeo struct {
	*Person
	Lovers map[string]*Person
}

func main() {

	boZhi := Person{
		Name: "Zhang Bo Zhi",
	}

	ziQi := Person{
		Name: "Li Zi Qi",
	}

	liHong := Romeo{
		Person: &Person{Name: "Wang Li Hong"},
		Lovers: make(map[string]*Person),
	}

	liHong.Lovers["boZhi"] = &boZhi
	liHong.Lovers["ziQi"] = &ziQi

	fmt.Println(liHong.Name + " has :")
	//output order is random
	for key, value := range liHong.Lovers {
		fmt.Println(key + ". " + value.Name)
	}
	/*
		PS C:\Users\D058009\go\helloWorldGo\maps> go run .\main.go
		Wang Li Hong has :
		ziQi. Li Zi Qi
		boZhi. Zhang Bo Zhi
		PS C:\Users\D058009\go\helloWorldGo\maps> go run .\main.go
		Wang Li Hong has :
		boZhi. Zhang Bo Zhi
		ziQi. Li Zi Qi
	*/
}
