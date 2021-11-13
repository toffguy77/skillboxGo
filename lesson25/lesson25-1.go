package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		str    string
		subStr string
	)
	err := parseInput(&str, &subStr)
	if err != nil {
		continue
	}
	match := searchByString(str, subStr)
	fmt.Println(match)
	match = searchByRune(str, subStr)
	fmt.Println(match)
}

func searchByRune(str string, str2 string) bool {
	return false
}

func searchByString(str string, str2 string) bool {
	return false
}

func parseInput(str *string, subStr *string) error {
	flag.StringVar(str, "string", "", "set string to search in")
	flag.StringVar(subStr, "string", "", "set string to search for matches")
	flag.Parse()

}
