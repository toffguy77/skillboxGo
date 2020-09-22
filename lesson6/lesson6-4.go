package main

import "fmt"

func main() {
	var (
		a, b, c int
		aMax    int = 10
		bMax    int = 100
		cMax    int = 1000
	)
	for i := 0; i < cMax; i++ {
		fmt.Println(a, b, c)
		c = c + 1
		if b >= bMax {
			continue
		}
		b = b + 1
		if a >= aMax {
			continue
		}
		a = a + 1
	}
	fmt.Println("Результат: ", a, b, c)

}
