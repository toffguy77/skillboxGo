package main

import "fmt"

func main() {
	var (
		studentsSum int = 10
		groupCnt    int = 3
		studentID   int
	)
	groups := make([][]int, groupCnt, groupCnt)

	for studentID := 1; studentID <= studentsSum; studentID++ {
		groupID := studentID % groupCnt
		groups[groupID] = append(groups[groupID], studentID)
	}
	fmt.Println("Студенты распределились по группам следующим образом: ", groups)
	fmt.Print("Введите порядковый номер студента: ")
	fmt.Scan(&studentID)
	fmt.Printf("Студент с порядковым номером %d находится в группе номер %d", studentID, studentID%groupCnt+1)
}
