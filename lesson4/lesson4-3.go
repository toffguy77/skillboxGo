package main

import "fmt"

func main() {
	const (
		MAX   int = 100000
		VALUE int = 100
	)
	var userVal int

	fmt.Print("Введиту сумму для выдачи: ")
	fmt.Scan(&userVal)

	if userVal > MAX {
		fmt.Println("Введена сумма, превышающая максимально допустимую в %d рублей", MAX)
		return
	}

	if userVal%VALUE == 0 {
		fmt.Println("Введена допустимая сумма для выдачи в банкомате")
	} else {
		fmt.Println("Введена недопустимая сумма для выдачи в банкомате")
	}
}
