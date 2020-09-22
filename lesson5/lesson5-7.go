package main

import "fmt"

func main() {
	var (
		number int
	)
	fmt.Print("Введите четырехзначное число: ")
	fmt.Scan(&number)
	if number < 1000 || number > 9999 {
		fmt.Println("Введено некоректное число")
		return
	}
	d := number % 10
	c := (number%100 - d) / 10
	b := (number%1000 - d - c) / 100
	a := (number%10000 - d - c - b) / 1000
	// fmt.Printf("Число разложено на разряды: %d %d %d %d\n", a, b, c, d)

	if a == d && b == c {
		fmt.Println("У вас зеркальный билетик!")
		return
	}

	if a+b == c+d {
		fmt.Println("У вас счастливый билетик!")
	} else {
		fmt.Println("У вас обычый билетик")
	}
}
