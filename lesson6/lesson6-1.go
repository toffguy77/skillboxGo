package main

import "fmt"

func main() {
	var (
		userInput int
	)

	fmt.Print("Введите целое положительное число: ")
	fmt.Scan(&userInput)

	for i := 0; i < userInput; i++ {
		fmt.Println(i)
	}
}
