package main2

import (
	"fmt"
	"strconv"
)

type Human struct {
	Name       string
	Body       string
	Job        string
	Aspiration string
	Love       *Human
}

type Smarty struct {
	*Human
	Level int
}

func (human *Human) print() string {
	var outputStr = "\nHi, I am " + human.Name + ".\nI am proud of my " + human.Body + " shape.\nI am currently a " + human.Job + " and I really love " + human.Aspiration

	if human.Love != nil {
		outputStr = outputStr + "\nI am in love with " + human.Love.Name
	}

	return outputStr

}
func (smarty *Smarty) print() string {
	var outputStr = smarty.Human.print()
	outputStr = outputStr + "\nI am a level " + strconv.Itoa(smarty.Level) + " smarty!"

	return outputStr
}

func main() {

	dummy := &Human{
		Name:       "Wang Li Hong",
		Body:       "Built",
		Job:        "Actor",
		Aspiration: "Dogs",
		Love:       nil,
	}

	prettySmart := &Smarty{
		&Human{
			"Zhang Bo Zhi",
			"Slim",
			"Actress",
			"Cats",
			dummy,
		},
		5,
	}

	fmt.Println(dummy.print())
	fmt.Println(prettySmart.print())
}
