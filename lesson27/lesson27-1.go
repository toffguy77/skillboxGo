package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Student struct {
	name  string
	age   int
	grade int
}
type StudentClass map[string]*Student

var (
	sc = StudentClass{}
)

func NewStudent(name string, age int, grade int) *Student {
	return &Student{name: name, age: age, grade: grade}
}

func (sc StudentClass) put(student *Student) error {
	if sc.get(student.name) != nil {
		return errors.New(fmt.Sprintf("Student %s already is in the class", student.name))
	}
	sc[student.name] = student
	log.Printf("[Added] %s : %+v\n", student.name, student)
	return nil
}

func (sc StudentClass) get(studentName string) *Student {
	if sc[studentName] != nil {
		return sc[studentName]
	}
	log.Printf("[Not Found] %s\n", studentName)
	return nil
}

func (sc StudentClass) form() {
	fmt.Println("Expected input: `<name> <age> <grade>`\nExample: Dima 38 18\nHit `Ctrl+Z` to stop input")
	for {
		var (
			name  string
			age   int
			grade int
		)

		_, err := fmt.Scanf("%s %d %d", &name, &age, &grade)
		if err != nil {
			log.Printf("[Error] %s\n", err)
		}
		err = sc.put(NewStudent(name, age, grade))
		if err != nil {
			log.Printf("[Error] %s\n", err)
		}
	}
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println()
		for _, student := range sc {
			fmt.Printf("%+v\n", sc.get(student.name))
		}
		os.Exit(0)
	}()

	sc.form()
}
