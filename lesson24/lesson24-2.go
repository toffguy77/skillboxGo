package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		n     int
		array []int
	)
	createArray(&n, &array)
	fmt.Printf("Initial array: %v\n", array)
	sortByBubbleDecs(&array)
	fmt.Printf("Sorted array: %v\n", array)
}

func sortByBubbleDecs(array *[]int) {
	if len(*array) == 1 {
		return
	}
	for i := 0; i < len(*array); i++ {
		for j := 1; j < len(*array)-i; j++ {
			if (*array)[j-1] < (*array)[j] {
				(*array)[j], (*array)[j-1] = (*array)[j-1], (*array)[j]
			}
		}
	}
}

func createArray(num *int, array *[]int) {
	for {
		fmt.Print("How many numbers there will in the array? ")
		_, err := fmt.Scanf("%d\n", num)
		if err != nil {
			fmt.Printf("error reading user input: %v\n", err)
			continue
		}
		if *num < 1 {
			fmt.Println("There should be one or more numbers")
		} else {
			break
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please start submitting sentences below")
	for i := 0; i < *num; i++ {
		var (
			userInput string
		)
		fmt.Printf("%d: ", i+1)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading user input: %v\n", err)
			i -= 1
			continue
		}
		userInputInt, err := strconv.Atoi(strings.TrimSpace(userInput))
		if err != nil {
			fmt.Println("There should be one or more numbers")
			i -= 1
			continue
		}
		*array = append(*array, userInputInt)
	}
}
