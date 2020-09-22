package main

import "fmt"

func main() {
	fmt.Print("Введите высоту елочки: ")
	var size int
	fmt.Scan(&size)

	minPlace := size - 1
	maxPlace := size - 1
	for i := 0; i < size; i++ {
		for j := 0; j < size*2-1; j++ {
			if j >= minPlace && j <= maxPlace {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		minPlace--
		maxPlace++
		fmt.Println()
	}
}
