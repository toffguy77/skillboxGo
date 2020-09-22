package main

import "fmt"

var (
	total        uint
	cost         uint
	deliveryCost uint
	discount     uint
)

func main() {
	fmt.Println("Пожалуйста, введите стоимость товара (целое положительное число):")
	fmt.Scan(&cost)
	fmt.Println("Теперь необходимо ввести стоимость доставки товара (целое положительное число):")
	fmt.Scan(&deliveryCost)
	fmt.Println("И, наконец, введите размер скидки, если она есть (целое положительное число):")
	fmt.Scan(&discount)

	total = cost + deliveryCost - discount

	fmt.Println("Полная стоимость товара составляет:", total)
}
