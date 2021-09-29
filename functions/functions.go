package functions

import (
	"fmt"
)

func StandardFunction() {
	fmt.Println("This is a standard function")
}

func FunctionWithParams(name string) {
	fmt.Printf("This is a function that has a parameter which is: %s\n", name)
}

func FunctionWithReturnType() string {
	return "This is a function with return type which is a string!"
}

func VariadicFunction(names ...string) {
	for _, name := range names {
		fmt.Println("This is from a variadic function: " + name)
	}
}

func RecursiveFunction(age int) bool {
	if age > 10 {
		return true
	}

	return RecursiveFunction(age - 1)
}

