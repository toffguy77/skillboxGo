package main

import (
	"fmt"
)

func main() {
	var (
		money      float64
		perc       float64
		years      int
		bankIncome float64
	)

	fmt.Print("Введите сумму для вклада: ")
	_, _ = fmt.Scan(&money)

	fmt.Print("Введите годовой процент: ")
	_, _ = fmt.Scan(&perc)

	fmt.Print("Введите количество лет для инвестиций: ")
	_, _ = fmt.Scan(&years)

	for i := 1; i <= 12*years; i++ {
		income := float64(money*perc) / 100
		bankIncome += (money + income)
		money += float64(int(income*100)) / 100
		bankIncome -= money
		fmt.Println(i, income, money, bankIncome)
	}
}
