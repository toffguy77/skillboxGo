package main

import "fmt"

const n = 4

func main() {
	nums := enterSlice()

	fmt.Printf("Исходный массив: %d\n", nums)
	switchSlice(&nums)
	fmt.Printf("Реверсивный массив: %d\n", nums)
}

func enterSlice() [n]int {
	var nums [n]int
	for i := 0; i < n; i++ {
		fmt.Printf("Введите %d элемент массива\n", i+1)
		fmt.Scan(&nums[i])
	}
	return nums
}

func switchSlice(nums *[n]int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
