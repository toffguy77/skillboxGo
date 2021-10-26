package main

import "fmt"

var (
	A = [3][5]int{
		{
			1, 2, 3, 4, 5,
		},
		{
			4, 5, 1, 2, 3,
		},
		{
			1, 5, 2, 3, 4,
		},
	}
	B = [5][4]int{
		{
			1, 2, 3, 5,
		},
		{
			4, 5, 1, 4,
		},
		{
			1, 5, 2, 3,
		},
		{
			1, 2, 3, 2,
		},
		{
			4, 5, 1, 1,
		},
	}
)

func main() {
	resArray := multiplyMatrix(A, B)
	fmt.Println(resArray)
}

func multiplyMatrix(A [3][5]int, B [5][4]int) [3][4]int {
	var C [3][4]int
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0])-1; j++ {
			C[i][j] = getMatrixValue(A, B, i, j)
		}
	}
	return C
}

func getMatrixValue(A [3][5]int, B [5][4]int, i int, j int) int {
	var res int
	for iter := 0; iter < len(A[0]); iter++ {
		res += A[i][iter] * B[iter][j]
	}
	return res
}
