package main

import "fmt"

func main() {
	var n int
	fmt.Print("please enter number of opening brackets N: ")
	fmt.Scan(&n)
	n = n * 2

	res := make([]string, n)
	var (
		cnt  int
		indx int
	)

	makeList(cnt, indx, n, res)
	fmt.Println("results:", res)
}

func makeList(cnt int, indx int, n int, res []string) {
	if cnt <= n-indx-2 {
		res[indx] = "("
		makeList(cnt+1, indx+1, n, res)
	}
	if cnt > 0 {
		res[indx] = ")"
		makeList(cnt-1, indx+1, n, res)
	}
	if indx == n && cnt == 0 {
		fmt.Println(res)
		return
	}

}
