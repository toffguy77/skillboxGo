package main

import "fmt"

var nums [10]int

func main() {
	createArray()
	odds, evens := calcOdds()
	fmt.Printf("Количечество четных %d, нечетных %d\n", evens, odds)
}

func calcOdds() (int, int) {
	var odds, evens int
	for _, i := range nums {
		if i%2 == 0 {
			evens++
			continue
		}
		odds++
	}
	return odds, evens
}

func createArray() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Введите %d элемент массива\n", i+1)
		fmt.Scan(&nums[i])
	}
}
