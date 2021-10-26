package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x int16   = 25
		y uint8   = 144
		z float32 = 2.3
	)

	fmt.Println(doCalc(x, y, z))
}

func doCalc(x int16, y uint8, z float32) float32 {
	s := float32(2*x) + float32(math.Pow(float64(y), 2)) - 3/z
	return s
}
