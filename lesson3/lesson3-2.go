package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		costDriver         float64 = 0
		costDriverRate     float64 = 0.25
		costFuel           float64 = 0
		costFuelRate       float64 = 0.2
		costTaxes          float64 = 0
		costTaxesRate      float64 = 0.2
		costRepair         float64 = 0
		costRepairRate     float64 = 0.2
		costTicket         float64 = 20
		countPassengers    float64 = 0
		countInPassengers  float64 = 0
		countOutPassengers float64 = 0
		totalIncome        float64 = 0
	)
	pointList := []string{"Первая", "Вторая", "Третья", "Четвертая"}

	for i, point := range pointList {
		fmt.Println("Прибываем на остановку "+point+". В салоне пассажиров:", countPassengers)
		if i == len(pointList)-1 {
			fmt.Println("---------Приехали--------")
			break
		}
		fmt.Println("Сколько пассажиров вышло на остановке?")
		fmt.Scan(&countOutPassengers)
		if (countPassengers - countOutPassengers) < 0 {
			fmt.Println("Из маршриутки не может выйти больше, чем есть внутри ;)")
			os.Exit(1)
		}
		fmt.Println("Сколько пассажиров зашло на остановке?")
		fmt.Scan(&countInPassengers)

		countPassengers = countPassengers + countInPassengers - countOutPassengers
		income := countPassengers * costTicket
		totalIncome += income
		costDriver += income * costDriverRate
		costFuel += income * costFuelRate
		costTaxes += income * costTaxesRate
		costRepair += income * costRepairRate

		fmt.Println("Отправляемся с остановки "+point+". В салоне пассажиров:", countPassengers)
		fmt.Println("-----------Едем---------")
	}

	total := totalIncome - (costDriver + costFuel + costTaxes + costRepair)

	fmt.Println("Всего заработали:", totalIncome, "руб.")
	fmt.Println("Зарплата водителя:", costDriver, "руб.")
	fmt.Println("Расходы на топливо:", costFuel, "руб.")
	fmt.Println("Налоги:", costTaxes, "руб.")
	fmt.Println("Расходы на ремонт машины:", costRepair, "руб.")
	fmt.Println("Итого доход:", total,"руб.")
}
