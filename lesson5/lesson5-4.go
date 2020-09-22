package main

import "fmt"

func main() {
	var (
		cost      int
		coinOne   int
		coinTwo   int
		coinThree int
	)
	fmt.Print("Введите стоимость товара: ")
	fmt.Scan(&cost)
	fmt.Println("Введите номиналы имеющихся монет")
	fmt.Print("первая монета: ")
	fmt.Scan(&coinOne)
	fmt.Print("вторая монета: ")
	fmt.Scan(&coinTwo)
	fmt.Print("третья монета: ")
	fmt.Scan(&coinThree)

	if coinOne == cost || coinTwo == cost || coinThree == cost || coinOne+coinTwo == cost || coinOne+coinThree == cost || coinTwo+coinThree == cost || coinOne+coinTwo+coinThree == cost {
		fmt.Println("Вы можете оплатить товар без сдачи")
	} else {
		fmt.Println("У вас не получится оплатить товар")
	}
}
