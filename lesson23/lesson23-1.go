package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var n int
	userInput(&n, "Please enter array lenght: ")
	array := createArray(n)
	fmt.Println("Original array", array)
	odds, evens := splitArray(array)
	fmt.Printf("Odds: %v\nEvens: %v\n", odds, evens)
}

func splitArray(array []int) ([]int, []int) {
	var (
		odds, evens []int
	)
	for _, v := range array {
		if v%2 != 0 {
			odds = append(odds, v)
		} else {
			evens = append(evens, v)
		}
	}
	return odds, evens
}

func userInput(n *int, text string) {
	fmt.Print(text)
	_, err := fmt.Scan(n)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func createArray(n int) (array []int) {
	rand.Seed(time.Now().UnixNano())
	array = make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(100)
	}
	return
}
