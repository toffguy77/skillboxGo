package main

import "fmt"

func main() {
	var (
		minNum int = 1
		maxNum int = 10
		answer string
	)
	fmt.Println("Загадайте целое число от 1 до 10 (включительно)")
	for i := 0; i < 4; i++ {
		midNum := (minNum + maxNum) / 2
		fmt.Printf("Я думаю, вы загадали %d. Скажите, верно ли это? если нет, то больше или меньше ваше загаданое число? ('=', '<', '>'): ", midNum)
		fmt.Scan(&answer)
		if answer == "=" {
			fmt.Println("Ура, я угадал!")
			return
		}
		if answer == ">" {
			minNum = midNum
		} else if answer == "<" {
			maxNum = midNum
		} else {
			fmt.Println("К сожалению, я не могу распознать ваш ответ...")
			return
		}
	}
	fmt.Println("Ну что ж, я не смог угадать ваше задуманное число. Поздравляю!")
}
