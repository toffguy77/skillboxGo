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
	privLuckyTicket := startPoint
	maxDistance := 0
	for i := startPoint; i <= endPoint; i++ {
		var sumA, sumB int
		for j := 0; j < digits; j++ {
			num := i / int(math.Pow10(j)) % 10
			if j < digits/2 {
				sumA += num
			} else {
				sumB += num
			}
		}
		if sumA == sumB {
			if i-privLuckyTicket > maxDistance {
				maxDistance = i - privLuckyTicket
			}
			privLuckyTicket = i
		}
	}
	fmt.Println(maxDistance)
}
