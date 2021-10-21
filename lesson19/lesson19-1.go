package main

import "fmt"

const (
	arrayLenOne = 4
	arrayLenTwo = 5
)

func main() {
	array1 := [arrayLenOne]int{1,3,5,6}
	array2:= [arrayLenTwo]int{2,3,7,8,9}

	fmt.Println(concatArrays(array1, array2))
}

func concatArrays(array1 [arrayLenOne]int, array2 [arrayLenTwo]int) [arrayLenOne+arrayLenTwo]int {
	 var (
		 i1 = 0
		 i2 = 0
		 resArray [arrayLenOne+arrayLenTwo] int
	 )
	 for i1+i2 < arrayLenOne+arrayLenTwo {
		 if i1 >= arrayLenOne {
			 resArray[i1+i2] = array2[i2]
			 i2+=1
			 continue
		 }
		 if i2 >= arrayLenTwo {
			 resArray[i1+i2] = array1[i1]
			 i1+=1
			 continue
		 }
		 if array1[i1]<array2[i2] {
			 resArray[i1+i2] = array1[i1]
			 i1 +=1
		 } else {
			 resArray[i1+i2] = array2[i2]
			 i2+=1
		 }
	 }
	 return resArray
}
