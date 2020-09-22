package main

import "fmt"

func main() {
	var (
		x int
		y int
	)
	fmt.Println("Введите координаты")
	fmt.Print("точка по Х оси: ")
	fmt.Scan(&x)
	fmt.Print("точка по Y оси: ")
	fmt.Scan(&y)

	if x == 0 && y == 0 {
		fmt.Printf("Точка ( %d, %d) находится в центре", x, y)
		return
	}

	if x > 0 {
		if y > 0 {
			fmt.Printf("Точка ( %d, %d) находится в I четверти", x, y)
		} else {
			fmt.Printf("Точка ( %d, %d) находится во II четверти", x, y)
		}
	} else {
		if y > 0 {
			fmt.Printf("Точка ( %d, %d) находится в IV четверти", x, y)
		} else {
			fmt.Printf("Точка ( %d, %d) находится во III четверти", x, y)
		}
	}
}
