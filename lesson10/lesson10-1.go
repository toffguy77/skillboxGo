package main

import (
	"fmt"
	"math"
)

/*
Разложение e^x в ряд Тейлора нашел тут https://dpva.ru/Guide/GuideMathematics/SeriesOfTaylorMaklorenFourier/SeriesOfTaylor/

Я не понимаю,как в ваших тестах для x=1 с разной точностьюполучаются разные результаты. Ведь при x=1dct 2+ члены рядапревращаются в 0.

И поэтому при x=1 результат всегда будет один, независимоот заданной точности и будет в точности совпадать с e.
*/

func main() {
	var (
		x          float64
		epsilonInt uint
	)

	fmt.Print("Введите значение Х для расчета e^x: ")
	_, _ = fmt.Scan(&x)

	fmt.Print("Введите желаемую точность вычислений (количество знаков после запятой): ")
	_, _ = fmt.Scan(&epsilonInt)

	epsilon := 1 / math.Pow(10, float64(epsilonInt))

	var (
		prevResult float64
		i          float64 = 1
	)

	result := math.E
	for math.Abs(prevResult-result) >= epsilon {
		prevResult = result
		result = prevResult + math.E*math.Pow(x-1, i)/factorial(i)
		i++
	}
	fmt.Println(result)
}

func factorial(x float64) float64 {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
