package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	initStr := "a10 10 20b 20 30c30 30 dd"
	strSlice := strings.Fields(initStr)

	for _, word := range strSlice {
		num, err := strconv.Atoi(word)
		if err == nil {
			fmt.Println(num)
		}
	}
}
