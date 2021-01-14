package main

import (
	"fmt"
	"strings"
)

func main() {
	textString := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	textSlices := strings.Fields(textString)

	var count int

	for _, word := range textSlices {
		if word[:1] >= "A" && word[:1] <= "Z" {
			count++
		}
	}

	fmt.Println(count)
}
