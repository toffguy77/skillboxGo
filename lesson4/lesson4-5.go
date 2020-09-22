package main

import "fmt"

func main() {
	var (
		day                    int
		guestsCnt              int
		cost                   int
		fridayDiscountMinLevel int = 10000
		fridayDiscount         int = 5
		mondayDiscount         int = 10
		manyPeopleFeeLevel     int = 5
		manyPeopleFee          int = 10
		grandTotal             float64
	)
	fmt.Print("Введите день недели, где \"1\" это понедельник, а \"7\" - воскресение: ")
	fmt.Scan(&day)
	fmt.Print("Сколько гостей? ")
	fmt.Scan(&guestsCnt)
	fmt.Print("Какова сумма по чеку? ")
	fmt.Scan(&cost)

	grandTotal = float64(cost)
	if day == 1 {
		grandTotal -= grandTotal * float64(mondayDiscount) / 100
	}
	if (day == 5) && (grandTotal > float64(fridayDiscountMinLevel)) {
		grandTotal -= grandTotal * float64(fridayDiscount) / 100
	}
	if guestsCnt > manyPeopleFeeLevel {
		grandTotal += grandTotal * float64(manyPeopleFee) / 100
	}

	fmt.Printf("Сумма к оплате составляет: %.2f рублей", grandTotal)
}
