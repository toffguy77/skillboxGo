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

	if x > 0 || y > 0 || z > 0 {
		fmt.Println("Среди введенных чисел есть одно (или более) больше 0")
	} else {
		fmt.Println("Среди введеных чисел нет положительных")
	}
}
