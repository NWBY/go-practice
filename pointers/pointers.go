package pointers

import (
	"fmt"
)

func PointersOutter() {
	rubysAge := 10
	passingPointers(&rubysAge)
}

func passingPointers(dogAge *int) {
	fmt.Printf("This is the pointer address in memory: %v\n", dogAge)
	fmt.Printf("This is the pointer value: %v\n", *dogAge)
}