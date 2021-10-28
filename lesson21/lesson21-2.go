package main

import "fmt"

func main() {
	runFunc(func(a int, b int) int { return a + b }, 2, 3)
	runFunc(func(a int, b int) int { return a - 2*b }, 3, 4)
	runFunc(func(a int, b int) int { return a * b }, 4, 5)
}

func runFunc(oper func(int, int) int, a int, b int) (res int) {
	defer func() {
		res = oper(a, b)
		fmt.Println(res)
	}()
	fmt.Printf("результат выполнения функции: ")
	res = 0
	return

}
