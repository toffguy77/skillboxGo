package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		f1  string
		f2  string
		res string
	)
	err := parseInput(&f1, &f2, &res)
	if err != nil {
		log.Fatalln(err)
	}
	err = mergeBytes(f1, f2, res)
	if err != nil {
		log.Fatalln("error merging []bytes files:", err)
	}
	err = mergeStrings(f1, f2, res)
	if err != nil {
		log.Fatalln("error merging []strings files:", err)
	}
}

func mergeStrings(f1 string, f2 string, res string) error {
	var (
		result []string
		err    error
	)
	if f1 != "" && f2 == "" {
		result, err = readFileLines(f1)
		if err != nil {
			return err
		}
	} else if f1 == "" && f2 != "" {
		result, err = readFileLines(f2)
		if err != nil {
			return err
		}
	} else {
		resultOne, err := readFileLines(f1)
		if err != nil {
			return err
		}
		resultTwo, err := readFileLines(f2)
		if err != nil {
			return err
		}
		result = append(resultOne, resultTwo...)
	}
	resultString := strings.Join(result, " ")
	if res == "" {
		fmt.Printf("Result for []strings is:\n%s\n", resultString)
	} else {
		err = saveResult([]byte(resultString), res)
		if err != nil {
			return fmt.Errorf("can't save results to the file %s: %v", res, err)
		}
	}
	return nil
}

func mergeBytes(f1 string, f2 string, res string) error {
	var (
		result []byte
		err    error
	)
	if f1 != "" && f2 == "" {
		result, err = readFileContent(f1)
		if err != nil {
			return err
		}
	} else if f1 == "" && f2 != "" {
		result, err = readFileContent(f2)
		if err != nil {
			return err
		}
	} else {
		resultOne, err := readFileContent(f1)
		if err != nil {
			return err
		}
		resultTwo, err := readFileContent(f2)
		if err != nil {
			return err
		}
		resultOne = append(resultOne, []byte(" ")...)
		result = append(resultOne, resultTwo...)
	}
	if res == "" {
		fmt.Printf("Result for []bytes is:\n%s\n", string(result))
	} else {
		err = saveResult(result, res)
		if err != nil {
			return fmt.Errorf("can't save results to the file %s: %v", res, err)
		}
	}
	return nil
}

func saveResult(result []byte, res string) error {
	file, err := os.Create(res)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(file)
	if err != nil {
		return fmt.Errorf("can't create file %s: %v", res, err)
	}
	_, err = file.Write(result)
	if err != nil {
		return fmt.Errorf("can't write to file %s: %v", res, err)
	}
	return nil
}

func readFileContent(file string) ([]byte, error) {
	result, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", file, err)
	}
	return result, nil
}

func readFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %v", filename, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text, nil
}

func parseInput(f1 *string, f2 *string, res *string) error {
	flag.StringVar(f1, "f1", "", "first file to merge")
	flag.StringVar(f2, "f2", "", "second file to merge")
	flag.StringVar(res, "res", "", "file to write results to")
	flag.Parse()
	if *f1 == "" && *f2 == "" {
		return fmt.Errorf("no files were provided for merge")
	}
	return nil
}
