package main

import "fmt"

func main() {
	var (
		duration    int
		orderTime   int
		packageTime int
	)

	fmt.Println("Эта программа рассчитает, сколько клиентов успеет обслужить кассир за смену.")
	fmt.Println("Введите длительность смены в минутах:")
	fmt.Scan(&duration)
	fmt.Println("Сколько минут клиент делает заказ?")
	fmt.Scan(&orderTime)
	fmt.Println("Сколько минут кассир собирает заказ?")
	fmt.Scan(&packageTime)
	fmt.Println("-----Считаем-----")

	clients := duration / (orderTime + packageTime)

	fmt.Println("За смену длиной", duration, "минут кассир успеет обслужить", clients, "клиентов.")
}
