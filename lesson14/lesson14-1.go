/*
Написать программу, которая на вход получается число. Далее число подается в функцию, которая возвращает true, если число четное, и false, если нечетное. Вывести в консоль результат функции.
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Print("please enter a number (int): ")
	fmt.Scan(&num)
	fmt.Printf("%d is even: %v\n", num, isEven(num))
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
