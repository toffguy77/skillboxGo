package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		maxUint8  int
		maxUint16 int
	)

	for i := 1; i < math.MaxUint32; i++ {
		if i&math.MaxUint8 == 0 {
			maxUint8++

		}
		if i&math.MaxUint16 == 0 {
			maxUint16++
		}
	}
	fmt.Printf("Количество переполнений: %d Uint8, %d Uint16", maxUint8, maxUint16)
}
