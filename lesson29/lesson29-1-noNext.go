package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Result struct {
	input  int
	square int
	double int
}

func main() {
	var (
		wg      sync.WaitGroup
		res     Result
		logfile string
	)

	flag.StringVar(&logfile, "log", "logfile", "output log filename")
	flag.Parse()

	toSquareChan := input(&wg, res)
	toDoubleChan := squared(&wg, toSquareChan)
	resChan := doubled(&wg, toDoubleChan)
	printResults(&wg, resChan, logfile)

	wg.Wait()
}

func input(wg *sync.WaitGroup, res Result) chan Result {
	toSquareChan := make(chan Result)
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("enter number: ")
			scannedItem, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("bad user input: %v\n", err)
				return
			}

			scannedItem = strings.Trim(scannedItem, "\n")
			if strings.ToLower(strings.Trim(scannedItem, " ")) == "stop" {
				break
			}

			val, err := strconv.Atoi(scannedItem)
			if err != nil {
				log.Printf("bad user input, not a number: %v: %v\n", scannedItem, err)
				continue
			}
			res.input = val
			toSquareChan <- res
		}
		log.Println("execution stopped successfully")
		os.Exit(0)
	}()

	return toSquareChan
}

func squared(wg *sync.WaitGroup, toSquareChan chan Result) chan Result {
	toDoubleChan := make(chan Result)
	wg.Add(1)
	go func() {
		defer close(toSquareChan)
		defer wg.Done()
		for res := range toSquareChan {
			res.square = res.input * res.input
			toDoubleChan <- res
		}

	}()
	return toDoubleChan
}

func doubled(wg *sync.WaitGroup, toDoubleChan chan Result) chan Result {
	resChan := make(chan Result)
	wg.Add(1)
	go func() {
		defer close(toDoubleChan)
		defer wg.Done()
		for res := range toDoubleChan {
			res.double = res.square * 2
			resChan <- res
		}
	}()
	return resChan
}

func printResults(wg *sync.WaitGroup, resChan chan Result, logfile string) {
	wg.Add(1)
	log.SetOutput(setLogFile(logfile))

	go func() {
		defer close(resChan)
		defer wg.Done()
		for res := range resChan {
			log.Printf("input: %d, squared: %d, doubled: %d", res.input, res.square, res.double)
		}
	}()
}

func setLogFile(logfile string) (f *os.File) {
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return f
}
