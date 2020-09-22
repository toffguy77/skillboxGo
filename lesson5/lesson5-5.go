package main

import "fmt"

func main() {
	var (
		stavkaOne   int
		stavkaTwo   int
		stavkaThree int
	)
	fmt.Println("Введите ставки по депозитам")
	fmt.Print("первая ставка: ")
	fmt.Scan(&stavkaOne)
	fmt.Print("вторая ставка: ")
	fmt.Scan(&stavkaTwo)
	fmt.Print("третья ставка: ")
	fmt.Scan(&stavkaThree)

	if stavkaOne < stavkaTwo && stavkaOne < stavkaThree {
		fmt.Printf("Самые высокие ставки это: %d и %d", stavkaTwo, stavkaThree)
	} else if stavkaTwo < stavkaOne && stavkaTwo < stavkaThree {
		fmt.Printf("Самые высокие ставки это: %d и %d", stavkaOne, stavkaThree)
	} else if stavkaThree < stavkaTwo && stavkaThree < stavkaOne {
		fmt.Printf("Самые высокие ставки это: %d и %d", stavkaOne, stavkaTwo)
	} else {
		fmt.Println("Выбрать 2 максимальные ставки не получается")
	}
}
