/*
Написать программу, которая на вход получает число, затем с помощью двух функции преобразует его. Первая умножает, а вторая прибавляет число. Использовать именованные возвращаемые значения.
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Print("please enter a number (int): ")
	fmt.Scan(&num)

	x1 := f1(num)
	fmt.Println("step one:", x1)
	x2 := f2(x1)
	fmt.Println("step two:", x2)
}

func f1(num int) (x int) {
	x = num * 2
	return
}

func f2(num int) (x int) {
	x = num + 3
	return
}
