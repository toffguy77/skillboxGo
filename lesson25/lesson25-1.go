package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	var (
		str    string
		subStr string
	)
	err := parseInput(&str, &subStr)
	if err != nil {
		log.Fatalf("error reading params: %v", err)
	}
	matchByStr := searchByString(str, subStr)
	fmt.Printf("Searching for %s substring in %s: %v\n", subStr, str, matchByStr)
	matchByRune := searchByRune(str, subStr)
	fmt.Printf("Searching for %s substring in %s using runes: %v\n", subStr, str, matchByRune)
}

func searchByRune(str string, subStr string) bool {
	strRunes := []rune(str)
	subStrRunes := []rune(subStr)
	for i := 0; i < len(strRunes); i++ {
		match := false
		for j := 0; j < len(subStrRunes); j++ {
			if subStrRunes[j] == strRunes[i+j] {
				match = true
			} else {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func searchByString(str string, subStr string) bool {
	return strings.Contains(str, subStr)
}

func parseInput(str *string, subStr *string) error {
	flag.StringVar(str, "str", "", "set string to search in")
	flag.StringVar(subStr, "sub", "", "set string to search for matches")
	flag.Parse()
	if *str == "" || *subStr == "" {
		return fmt.Errorf("empty input params")
	} else {
		return nil
	}
}
