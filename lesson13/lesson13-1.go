package main

import "fmt"

func main() {
	s := "0123456789"
	initFunc(s, f1, f2)
}

func initFunc(s string, f1 func(string), f2 func(string)) {
	f2(s)
	f1(s)
}

func f1(s string) {
	for _, w := range s {
		fmt.Print(string(w))
	}
	fmt.Println()
}

func f2(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
	fmt.Println()
}
