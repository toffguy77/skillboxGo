package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./lesson12/lesson12-1.txt")
	if err != nil {
		fmt.Println("error: cant read from file:", err)
		return
	}

	fmt.Println(string(content))
}
