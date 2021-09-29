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