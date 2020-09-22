package main

import "fmt"

func main() {
	var (
		backetOne      int
		backetTwo      int
		backetThree    int
		backetOneMax   int
		backetTwoMax   int
		backetThreeMax int
	)

	fmt.Print("Введите размер корзины 1: ")
	fmt.Scan(&backetOneMax)
	fmt.Print("Введите размер корзины 2: ")
	fmt.Scan(&backetTwoMax)
	fmt.Print("Введите размер корзины 3: ")
	fmt.Scan(&backetThreeMax)

	i := 0
	for {
		if i < backetOneMax {
			backetOne++
		}
		if i < backetTwoMax {
			backetTwo++
		}
		if i < backetThreeMax {
			backetThree++
		}
		if i >= backetOneMax && i >= backetTwoMax && i >= backetThreeMax {
			break
		}
		fmt.Println("Итерация", i)
		fmt.Printf("Корзина #1: %d из %d\n", backetOne, backetOneMax)
		fmt.Printf("Корзина #2: %d из %d\n", backetTwo, backetTwoMax)
		fmt.Printf("Корзина #3: %d из %d\n", backetThree, backetThreeMax)
		i++
	}
}
