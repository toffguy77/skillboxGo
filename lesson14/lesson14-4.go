/*
Написать программу, в которой будет 3 функции, попарно использующие разные глобальные переменные. Функции, должны прибавлять к поданном на вход числу глобальную переменную и возвращать результат. Затем вызвать по очереди три функции, передавая результат из одной в другую.
*/

package main

import "fmt"

const (
	a = 5
	b = 4
	c = 3
)

func main() {
	z := 1
	y := f3(f2(f1(z)))
	fmt.Println("result:", y)
}

func f1(x int) int {
	fmt.Println("f1 result:", a+x)
	return a + x
}

func f2(x int) int {
	fmt.Println("f2 result:", b+x)
	return b + x
}

func f3(x int) int {
	fmt.Println("f3 result:", c+x)
	return c + x
}
