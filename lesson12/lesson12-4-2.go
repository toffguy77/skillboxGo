package main

import (
	"fmt"
	"io/ioutil"
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

	if err = ioutil.WriteFile(fileName, []byte{65, 66, 67}, 0644); err != nil {
		fmt.Println("error: cant write to file:", err)
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

	if err = ioutil.WriteFile(fileName, []byte{68, 69, 70}, 0644); err != nil {
		fmt.Println("error: cant write to file:", err)
		os.Chmod(fileName, 06444)
		return
	}

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error: cant read from file:", err)
		return
	}
	fmt.Println(string(content))
	fmt.Println("all passed")
}
