package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		n         int
		sentences []string
		letters   []rune
	)
	createNewArray(&n)
	readSentences(n, &sentences)
	createRuneArray(&letters)
	fmt.Printf("\nThe following []rune slice will be used %v which is `%v` in letters\n", letters, string(letters))
	lastWords := getLastWords(sentences)
	resArray := countLetters(letters, lastWords)
	fmt.Println(resArray)
}

func getLastWords(sentences []string) []string {
	var res []string
	for _, sentence := range sentences {
		word := strings.Fields(sentence)
		res = append(res, word[len(word)-1])
	}
	return res
}

func createRuneArray(letters *[]rune) {
	fmt.Print("Please start submitting case sensitive letters to check: ")
	reader := bufio.NewReader(os.Stdin)
	tempLetters, _ := reader.ReadString('\n')
	tempLetters = strings.TrimSpace(tempLetters)
	runeTempLetters := []rune(tempLetters)
	for _, tempLetter := range runeTempLetters {
		uniq := true
		for _, letter := range *letters {
			if tempLetter == letter {
				uniq = false
				break
			}
		}
		if uniq {
			*letters = append(*letters, tempLetter)
		}
	}
}

func countLetters(letters []rune, lastWords []string) [][]int {
	numsWords := len(lastWords)
	numsLetters := len(letters)
	res := createMatrix(numsWords, numsLetters)
	for i := 0; i < numsWords; i++ {
		for j := 0; j < numsLetters; j++ {
			res[i][j] += strings.IndexRune(lastWords[j], letters[i])
		}
	}
	return res
}

func createMatrix(words int, letters int) [][]int {
	matrix := make([][]int, words)
	for i := 0; i < words; i++ {
		matrix[i] = make([]int, letters)
	}
	return matrix
}

func createNewArray(num *int) {
	for {
		fmt.Print("How many sentences there will be? ")
		fmt.Scan(num)
		if *num < 1 {
			fmt.Println("There should be one or more sentences")
		} else {
			break
		}
	}
}

func readSentences(n int, sentences *[]string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please start submitting sentences below")
	for i := 0; i < n; i++ {
		fmt.Printf("%d: ", i+1)
		sentence, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading: %v", err)
			continue
		}
		sentence = strings.TrimSpace(sentence)
		if sentence == "" {
			fmt.Println("Sentence should contain at least one word")
			i -= 1
			continue
		}
		*sentences = append(*sentences, sentence)
	}
}
