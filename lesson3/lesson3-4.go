package main

import "fmt"

func main() {
	var (
		height           int = 100
		speedBamboo      int = 50
		speedCaterpillar int = 20
		finishHeight     int = 300
		countDays        int = 0
	)

	heightThirdDay := height + (speedBamboo-speedCaterpillar)*2 + speedBamboo/2
	fmt.Println("К середине третьего дня бамбук быдет высотой", heightThirdDay, "сантиметров")

	countDays = (finishHeight-height-speedBamboo)/(speedBamboo-speedCaterpillar) + 1
	fmt.Println("Бамбук выростет выше", finishHeight, "сантиметров на", countDays, "день")
}
