package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Введите месяц: ")
	var month string
	fmt.Scan(&month)

	switch strings.ToLower(month) {
	case "декабрь", "январь", "февраль":
		fmt.Println("зима")
	case "март", "аперль", "май":
		fmt.Println("весна")
	case "июнь", "июль", "август":
		fmt.Println("лето")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("осень")
	default:
		fmt.Println("это вообще не месяц")
	}
}
