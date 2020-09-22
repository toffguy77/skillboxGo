package main

import "fmt"

func main() {
	var (
		total          int
		entranceNumber int
		appNumber      int
	)

	fmt.Println("Сумма, указанная в квитанции:")
	fmt.Scan(&total)
	fmt.Println("Подъездов в доме:")
	fmt.Scan(&entranceNumber)
	fmt.Println("Квартир в каждом подъезде:")
	fmt.Scan(&appNumber)

	cheque := total / (entranceNumber * appNumber)

	fmt.Println("----Результат-----")
	fmt.Println("Каждая квартира должна заплатить по", cheque, "руб.")
}
