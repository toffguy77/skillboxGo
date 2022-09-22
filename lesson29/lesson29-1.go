package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	nextChan := make(chan bool, 1)
	nextChan <- true

	toSquareChan := input(&wg, nextChan)
	toDoubleChan := squared(&wg, toSquareChan)
	_ = doubled(&wg, toDoubleChan, nextChan)

	wg.Wait()
}

func input(wg *sync.WaitGroup, nextChan chan bool) chan int {
	toSquareChan := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("closing nextChan")
			close(nextChan)
		}()
		defer wg.Done()

		for <-nextChan {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("enter number: ")
			scannedItem, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("bad user input: %v\n", err)
				return
			}

			scannedItem = strings.Trim(scannedItem, "\n")
			if strings.ToLower(strings.Trim(scannedItem, " ")) == "stop" {
				break
			}

			val, err := strconv.Atoi(scannedItem)
			if err != nil {
				fmt.Printf("bad user input, not a number: %v: %v\n", scannedItem, err)
				nextChan <- true
				fmt.Println("------")
				continue
			}
			toSquareChan <- val
		}
		fmt.Println("execution stopped successfully")
		os.Exit(0)
	}()

	return toSquareChan
}

func squared(wg *sync.WaitGroup, toSquareChan chan int) chan int {
	toDoubleChan := make(chan int)
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("closing toSquareChan")
			close(toSquareChan)
		}()
		defer wg.Done()
		for val := range toSquareChan {
			val *= val
			fmt.Printf("squared: %d\n", val)
			toDoubleChan <- val
		}

	}()
	return toDoubleChan
}

func doubled(wg *sync.WaitGroup, toDoubleChan chan int, nextChan chan bool) chan bool {
	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("closing toDoubleChan")
			close(toDoubleChan)
		}()
		defer wg.Done()
		for val := range toDoubleChan {
			val = val * 2
			fmt.Printf("doubled: %d\n", val)
			fmt.Println("------")
			nextChan <- true
		}
	}()
	return nextChan
}
