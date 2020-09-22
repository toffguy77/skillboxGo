package main

import "fmt"

func main() {
	var (
		cost            int
		discountPerc    int
		discount        int
		maxDiscountPerc int = 30
		maxDiscount     int = 2000
	)
	fmt.Print("Введите стоимость товара: ")
	fmt.Scan(&cost)
	fmt.Print("Введите размер скидки в процентах: ")
	fmt.Scan(&discountPerc)

	if discountPerc > maxDiscountPerc {
		fmt.Printf("Размер введеной скридки превышает максимально допустимую. Будет применена максимально возможная скидка в %d процентов\n", maxDiscountPerc)
		discountPerc = maxDiscountPerc
	}

	discount = cost * discountPerc / 100
	if cost*discount/100 > maxDiscount {
		fmt.Printf("Размер рассчитанной скридки превышает максимально допустимую. Будет применена максимально возможная скидка в %d рублей\n", maxDiscount)
		discount = maxDiscount
	}

	userCost := cost - discount
	fmt.Printf("Стоимость товара: %d рублей. Ваша скидка составила: %d рублей. Итого к оплате: %d рублей", cost, discount, userCost)

}
