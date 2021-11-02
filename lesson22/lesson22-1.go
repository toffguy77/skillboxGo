package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var (
	n   int
	num int
)

func main() {
	userInput(&n, "Please enter array lenght: ")
	array := createArray(n)
	fmt.Println(array)
	userInput(&num, "Please enter a number to search for: ")
	rest := checkNum(array, num)
	if rest == 0 {
		fmt.Printf("The number %d was not found in the array\n", num)
	} else {
		fmt.Printf("The count for left numbers in the array is: %d\n", rest)
	}
}

func userInput(n *int, text string) {
	fmt.Print(text)
	_, err := fmt.Scan(n)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func checkNum(array []int, num int) int {
	pos := 0
	for i := 0; i < len(array); i++ {
		if array[i] == num {
			pos = i
			return len(array) - pos - 1
		}
	}
	return pos
}

func createArray(n int) (array []int) {
	rand.Seed(time.Now().UnixNano())
	array = make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(10)
	}
	return
}
