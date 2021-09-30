package control

import (
	"fmt"
)

func IfStatements(yearFounded uint) {
	if yearFounded == 2020 {
		fmt.Println("The business was founded in the year 2020!")
	} else if yearFounded == 2021 {
		fmt.Println("The business was founded in the year 2021!")
	} else {
		fmt.Println("The business was founded in the year unknown!")
	}
}

func SwitchStatements(name string) {
	switch name {
	case "Sam":
		fmt.Println("Switched on Sam!")
	case "Courtney":
		fmt.Println("Switched on Courtney!")
	default:
		fmt.Println("Switched on " + name + "!")
	}
}

func LoopMeBaby() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}
