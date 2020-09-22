package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a, b, c int
	)
	fmt.Println("Введите коэффициенты квадратного уравнения: a*x^2 + b*y + c = 0")
	fmt.Print("коэффициент a: ")
	fmt.Scan(&a)
	fmt.Print("коэффициент b: ")
	fmt.Scan(&b)
	fmt.Print("коэффициент c: ")
	fmt.Scan(&c)
	if c > 0 {
		fmt.Printf("Квадратное уравнение выглядит как: %d*x^2 + %d*y + %d = 0\n", a, b, c)
	} else if c < 0 {
		fmt.Printf("Квадратное уравнение выглядит как: %d*x^2 + %d*y - %d = 0\n", a, b, c)
	} else {
		fmt.Printf("Квадратное уравнение выглядит как: %d*x^2 + %d*y = 0\n", a, b)
	}

	d := b*b - 4*a*c
	fmt.Printf("Дискриминант для вашего уравнения: %d\n", d)

	if d > 0 {
		x1 := (-float64(b) + math.Sqrt(float64(d))) / (2 * float64(a))
		x2 := (-float64(b) - math.Sqrt(float64(d))) / (2 * float64(a))
		fmt.Printf("Корни вашего уравнения: %v, %v", x1, x2)
	} else if d == 0 {
		x := (-b) / (2 * a)
		fmt.Printf("Корень вашего уравнения: %v", x)
	} else {
		fmt.Println("Корней нет")
	}
}
