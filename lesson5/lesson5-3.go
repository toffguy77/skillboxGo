package main

import "fmt"

func main() {
	var (
		x, y, z int
	)
	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)

	fmt.Print("Введите третье число: ")
	fmt.Scan(&z)

	if z == y || x == z || y == z {
		fmt.Println("Есть совпадающие числа")
	} else {
		fmt.Println("Совпадающих чисел нет")
	}
}
