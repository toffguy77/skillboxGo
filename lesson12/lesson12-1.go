package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	const (
		layoutTime = "2006-01-02 15:04:05"
	)

	file, err := os.Create("lesson12-1.txt")
	if err != nil {
		fmt.Println("error creating new file:", err)
		return
	}
	defer file.Close()

	var (
		cnt   int
		input string
	)

	for {
		fmt.Print("plz submit your text: ")
		fmt.Scanln(&input)
		if strings.ToLower(input) == "exit" || strings.ToLower(input) == "выход" {
			fmt.Println("finishing the program...")
			return
		}
		cnt++
		if _, err = file.WriteString(fmt.Sprintf("%d %v %v\n", cnt, time.Now().Format(layoutTime), input)); err != nil {
			fmt.Println("error writing to the file:", err)
		}
	}
}
