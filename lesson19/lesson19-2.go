package main

import "fmt"

func main() {
	array := [6]int{9,8,7,6,5,4}
	fmt.Println(bubbleSort(array))
}

func bubbleSort(array [6]int) [6]int {
	for i:=0;i<len(array); i++{
		for j:=0;j<len(array)-1;j++{
			if array[j]>array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
