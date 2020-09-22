package main

import "fmt"

func main() {
	const acceptGrade int = 275

	exams := []string{"math", "physics", "literatire"}
	var studentGrade int
	for _, exam := range exams {
		var grade int
		fmt.Println("Введите вашу оценку по предмету", exam)
		fmt.Scan(&grade)
		studentGrade += grade
	}
	fmt.Println("Ваш уровень:", studentGrade)
	if studentGrade >= acceptGrade {
		fmt.Println("Ваш уровень достаточен для зачисления. Поздравляем!")
	} else {
		fmt.Println("К сожалению вашего уровня пока не достаточно для зачисления.")
	}
}
