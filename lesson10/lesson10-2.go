package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a   int16
		b   int16
		res int32
	)

	fmt.Print("Введите числа для перемножения: ")
	_, _ = fmt.Scan(&a, &b)
	fmt.Println(a, b)
	res = int32(a) * int32(b)
	if res >= 0 {
		calcForPos(res)
	} else {
		calcForNeg(res)
	}
}

func calcForPos(r int32) {
	switch {
	case r <= int32(math.MaxUint8):
		{
			fmt.Printf("Число %d помещается в тип Uint8\n", uint8(r))
		}
	case r <= int32(math.MaxUint16):
		{
			fmt.Printf("Число %d помещается в тип Uint16\n", uint16(r))
		}
	case r <= int32(math.MaxInt32):
		{
			fmt.Printf("Число %d помещается в тип Int32\n", int32(r))
		}
	default:
		{
			fmt.Printf("Число %d не вместилось в Uint32\n", uint32(r))
		}
	}
}

func calcForNeg(r int32) {
	switch {
	case r >= int32(math.MinInt8):
		{
			fmt.Printf("Число %d помещается в тип Int8\n", int8(r))
		}
	case r >= int32(math.MinInt16):
		{
			fmt.Printf("Число %d помещается в тип Int16\n", int16(r))
		}
	case r >= int32(math.MinInt32):
		{
			fmt.Printf("Число %d помещается в тип Int32\n", int32(r))
		}
	default:
		{
			fmt.Printf("Число %d не вместилось в Int32\n", int32(r))
		}
	}
}
