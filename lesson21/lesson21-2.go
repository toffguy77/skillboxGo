package main

import "fmt"

func main() {
	runFunc(func(a int, b int) int { return a + b })
	runFunc(func(a int, b int) int { return a - 2*b })
	runFunc(func(a int, b int) int { return a * b })
}

func runFunc(A func(int, int) int) {
	defer func() {
		res := A(2, 4)
		fmt.Println(res)
	}()
	fmt.Printf("результат выполнения функции: ")
}
