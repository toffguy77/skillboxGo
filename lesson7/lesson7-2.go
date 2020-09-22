package main

import "fmt"

func main() {
	var (
		iBlock string
		jBlock string
	)
	fmt.Print("Введите размерность шахматной дости: ")
	var size int
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			iBlock = " "
		} else {
			iBlock = "*"
		}
		for j := 0; j < size; j++ {
			if iBlock == " " {
				jBlock = "*"
			} else {
				jBlock = " "
			}
			if j%2 == 0 {
				fmt.Print(iBlock)
			} else {
				fmt.Print(jBlock)
			}
		}
		fmt.Println()
	}
}
