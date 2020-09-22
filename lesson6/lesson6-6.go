package main

import "fmt"

func main() {
	var (
		elevatorCapacity int  = 2
		needToCatch      bool = true
	)
	peopleOnFloor := map[int]bool{
		24: false,
		23: false,
		22: false,
		21: false,
		20: false,
		19: false,
		18: false,
		17: false,
		16: false,
		15: false,
		14: false,
		13: false,
		12: false,
		11: false,
		10: true,
		9:  false,
		8:  false,
		7:  true,
		6:  false,
		5:  false,
		4:  true,
		3:  false,
		2:  false,
		1:  false,
	}
	for needToCatch {
		for i := 1; i < len(peopleOnFloor); i++ {
			fmt.Printf("Лифт на %d этаже\n", i)
		}
		peopleInElevator := 0
		for i := len(peopleOnFloor); i > 0; i-- {
			if peopleInElevator < elevatorCapacity && i == 1 {
				needToCatch = false
			}
			fmt.Printf("Лифт на %d этаже, людей в лифте: %d\n", i, peopleInElevator)
			if !peopleOnFloor[i] {
				continue
			}
			if peopleInElevator >= elevatorCapacity {
				continue
			}
			fmt.Println("В лифт зашел 1 человек")
			peopleInElevator++
			peopleOnFloor[i] = false
		}
	}
}
