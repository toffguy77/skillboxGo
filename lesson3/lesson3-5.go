package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a, b = b, a

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
