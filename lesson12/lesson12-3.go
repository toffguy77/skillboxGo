package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./lesson12/lesson12-3.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error: cant create file:", err)
		return
	}
	defer file.Close()

	if _, err = file.WriteString("test message\n"); err != nil {
		fmt.Println("error: cant write to file:", err)
		return
	}

	buf, err := readFromFile(fileName)
	if err != nil {
		fmt.Println("error: cant get file stats:", err)
		return
	}
	file.Chmod(0444)
	file.Close()

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("error: cant create file:", err)
		return
	}
	defer file.Close()

	file.Sync()
	if _, err = file.WriteString("another test message\n"); err != nil {
		fmt.Println("error: cant write to file:", err)
		file.Chmod(0644)
		return
	}

	buf, err = readFromFile(fileName)
	if err != nil {
		fmt.Println("error: cant get file stats:", err)
		return
	}
	fmt.Println(string(buf))

	file.Chmod(0644)
	fmt.Println("all passed")
}

func readFromFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := stat.Size()
	buf := make([]byte, size)
	if _, err := io.ReadFull(file, buf); err != nil {
		return nil, err
	}
	return buf, nil
}
