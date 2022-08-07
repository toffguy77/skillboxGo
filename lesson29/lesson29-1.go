package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		firstChan := input()
		secondChan := squared(firstChan)
		_ = doubled(secondChan)
	}
}

func input() chan int {
	firstChan := make(chan int)
	go func() {
		defer close(firstChan)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("enter number: ")
		scannedItem, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("bad user input: %v\n", err)
			return
		}
		scannedItem = strings.Trim(scannedItem, "\n")
		if strings.ToLower(strings.Trim(scannedItem, " ")) == "stop" {
			fmt.Println("execution stopped successfully")

			os.Exit(0)
		}
		val, err := strconv.Atoi(scannedItem)
		if err != nil {
			fmt.Printf("bad user input, not a number: %v: %v\n", scannedItem, err)
			return
		}
		firstChan <- val
	}()
	return firstChan
}

func squared(firstChan <-chan int) chan int {
	val := <-firstChan
	secondChan := make(chan int)
	val *= val
	fmt.Printf("squared: %d\n", val)
	go func() {
		secondChan <- val
	}()
	return secondChan
}

func doubled(secondChan <-chan int) chan int {
	val := <-secondChan
	thirdChan := make(chan int)
	val = val * 2
	fmt.Printf("doubled: %d\n", val)
	fmt.Println("------")
	go func() {
		thirdChan <- val
	}()
	return thirdChan
}
