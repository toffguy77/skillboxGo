package main

import (
	"fmt"
)

func main() {
	a := 42
	b := 153

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	temp := b

	b = a
	a = temp

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
