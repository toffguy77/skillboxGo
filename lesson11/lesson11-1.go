package main

import (
	"fmt"
	"strings"
)

func main() {
	textString := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	extension := " Go это Язык программирования с Открытым исходным кодом, который позволяет Легко создавать простое, надежное и эффективное Программное обеспечение"
	textString = textString + extension
	textSlices := strings.Fields(textString)

	var count int

	for _, word := range textSlices {
		for _, r := range word {
			if r >= 'A' && r <= 'Z' {
				count++
			}
			if r >= 'А' && r <= 'Я' {
				count++
			}
		}
	}

	fmt.Println(count)
}
