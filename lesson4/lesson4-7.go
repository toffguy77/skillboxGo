package main

import "fmt"

func main() {
	const (
		level0fee float64 = 0.13
		level1            = 10000
		level1fee float64 = 0.2
		level2            = 50000
		level2fee float64 = 0.3
	)
	var (
		sum  int = 30000
		fees float64
	)
	//fmt.Print("Введите сумму, подлежащую для расчета налогообложения: ")
	//fmt.Scan(&sum)
	if sum < level1 {
		fees = float64(sum) * level0fee
	} else if sum < level2 {
		fees = level1*level0fee + float64(sum-level1)*level1fee
	} else {
		fees = level1*level0fee + (level2-level1)*level1fee + float64(sum-level2)*level2fee
	}

	fmt.Printf("Сумма налога, подлежащего уплате, составляет %.2f рублей", fees)

}
