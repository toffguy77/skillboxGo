package main

import "fmt"

func main() {
	var (
		a, b int
	)
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	sum := a + b

	for i := a; i < sum; i = i + 1 {
		a = a + 1
		fmt.Printf("a: %d, b: %d, sum: %d\n", a, b, sum)
	}
}
