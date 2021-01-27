package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./lesson12/lesson12-1.txt")
	if err != nil {
		fmt.Println("error: cant open the file:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("error: cant calc file size:", err)
		return
	}

	size := stat.Size()
	buf := make([]byte, size)
	if _, err = io.ReadFull(file, buf); err != nil {
		fmt.Println("error: cant read from file:", err)
		return
	}
	fmt.Println(string(buf))
}
