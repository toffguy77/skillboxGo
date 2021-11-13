package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		f1  string
		f2  string
		res string
	)
	parseInput(&f1, &f2, &res)
	err := merge(f1, f2, res)
	if err != nil {
		log.Fatalln("error merging files: ", err)
	}
}

func merge(f1 string, f2 string, res string) error {
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
		resultOne = append(resultOne, []byte("\n")...)
		result = append(resultOne, resultTwo...)
	}
	if res == "" {
		fmt.Printf("Result is:\n%s\n", string(result))
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

func parseInput(f1 *string, f2 *string, res *string) {
	flag.StringVar(f1, "f1", "", "first file to merge")
	flag.StringVar(f2, "f2", "", "second file to merge")
	flag.StringVar(res, "res", "", "file to write results to")
	flag.Parse()
}
