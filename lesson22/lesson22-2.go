package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var n int
	userNumInput(&n, "Please enter array length: ")
	array := createNewArray(n)
	fmt.Println(array)
	var num int
	userNumInput(&num, "Please enter a number to search for: ")
	index := findFirst(array, num)
	if index == -1 {
		fmt.Printf("The number %d was not found in the array\n", num)
	} else {
		fmt.Printf("The first occurence of %d is: %d\n", num, index+1)
	}
}

func findFirst(array []int, num int) int {
	var (
		startPoint = 0
		lastPoint  = len(array) - 1
		checkPoint = -1
	)
	for {
		pointToCheck := (startPoint + lastPoint) / 2
		if pointToCheck == len(array)-1 && array[pointToCheck] != num {
			break
		}
		if array[pointToCheck] > num {
			lastPoint = pointToCheck - 1
		} else if array[pointToCheck] < num {
			startPoint = pointToCheck + 1
		} else {
			checkPoint = pointToCheck
			break
		}
	}
	if checkPoint == -1 || checkPoint == 0 {
		return checkPoint
	} else {
		for {
			if array[checkPoint] == array[checkPoint-1] {
				checkPoint -= 1
			} else {
				break
			}
			if checkPoint == 0 {
				break
			}
		}
	}
	return checkPoint
}

func userNumInput(n *int, text string) {
	fmt.Print(text)
	_, err := fmt.Scan(n)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

func createNewArray(n int) (array []int) {
	rand.Seed(time.Now().UnixNano())
	array = make([]int, n)
	array[0] = rand.Intn(2)
	for i := 1; i < n; i++ {
		array[i] = array[i-1] + rand.Intn(2)
	}
	return
}
