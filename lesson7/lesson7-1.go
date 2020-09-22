package main

import (
	"fmt"
	"math"
)

var (
	startPoint int = 100000
	endPoint   int = 999999
	digits     int = 6
)

func main() {
	cnt := 0
OutterLoop:
	for i := startPoint; i <= endPoint; i++ {
		for j := 0; j < digits/2; j++ {
			a := i / int(math.Pow10(j)) % 10
			b := i / int(math.Pow10(digits-1-j)) % 10
			if a != b {
				continue OutterLoop
			}
		}
		cnt++
	}
	fmt.Println(cnt)
}
