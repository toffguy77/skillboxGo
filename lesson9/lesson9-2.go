package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Введите день недели: ")
	var day string
	fmt.Scan(&day)
	day = strings.ToLower(day)

	days := map[string]int{
		"пн": 1,
		"вт": 2,
		"ср": 3,
		"чт": 4,
		"пт": 5,
	}

	switch {
	case days[day] == 0:
		fmt.Println("это вообще не день из рабочей недели")
	case days[day] <= 1:
		fmt.Println("понедельник")
		fallthrough
	case days[day] <= 2:
		fmt.Println("вторник")
		fallthrough
	case days[day] <= 3:
		fmt.Println("среда")
		fallthrough
	case days[day] <= 4:
		fmt.Println("четверг")
		fallthrough
	case days[day] <= 5:
		fmt.Println("пятница")
	}
}
