/*
Написать программу, которая с помощью функции генерирует 3 случайные точки в двумерном пространстве (две координаты), а затем с помощью другой функции преобразует эти координаты по формулам: x1 = 2 * x + 10, y1 = -3 * y - 5.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := generateInt(10)
	fmt.Println(a, b)
	x, y := newValues(a, b)
	fmt.Println(x, y)
}

func generateInt(a int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(a), rand.Intn(a)
}

func newValues(a, b int) (x, y int) {
	x = 2*a + 10
	y = -3*b - 5
	return
}
