package main

import (
	"fmt"
	"log"
)

func main() {
	var array = [][]int{
		{
			2, 4, 1,
		},
		{
			0, 2, 1,
		},
		{
			2, 1, 1,
		},

	}
	det := calcDet(array)
	fmt.Println(det)
}

func calcDet(array [][]int) int {
	if len(array) != 3 {
		log.Fatalln("cannot calculate determinant for not 3x3 matrix")
	}
	det := array[0][0]*(array[1][1]*array[2][2]-array[1][2]*array[2][1])-
		array[0][1]*(array[1][0]*array[2][2]-array[1][2]*array[2][0])+
		array[0][2]*(array[1][0]*array[2][1]-array[1][1]*array[2][0])

	return det
}
