package main

import "fmt"

func main() {
	const (
		LENGHT int = 3
		MAX    int = 5
	)

	numbers := make([]int, LENGHT, LENGHT)
	for i := 0; i < LENGHT; i++ {
		fmt.Printf("Введите %d число: ", i+1)
		fmt.Scan(&numbers[i])
	}
	fmt.Println("Введеные данные: ", numbers)
	for _, num := range numbers {
		if num > MAX {
			fmt.Printf("Найдено число больше %d - это %d", MAX, num)
			return
		}
	}
	fmt.Printf("Числа больше %d не найдено", MAX)
}
