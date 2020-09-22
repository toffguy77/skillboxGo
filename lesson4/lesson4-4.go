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

	var cnt int
	for _, num := range numbers {
		if num > MAX {
			cnt++
		}
	}

	if cnt > 0 {
		if cnt == 1 {
			fmt.Printf("Найдено %d число больше %d", cnt, MAX)
		} else {
			fmt.Printf("Найдено %d числа больше %d", cnt, MAX)
		}

	} else {
		fmt.Printf("Чисел больше %d не найдено", MAX)
	}

}
