package main

import (
	"fmt"
	"lesson27/students"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	studDB := make(map[string]students.Student)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println()
		for k, v := range studDB {
			fmt.Printf("Name: %s, data: %+v\n", k, v)
		}
		os.Exit(0)
	}()

	err := formDB(&studDB)
	if err != nil {
		log.Fatal(err)
	}
}

func formDB(studDB *map[string]students.Student) error {
	fmt.Println("Expected input: `<name> <age> <grade>`\nExample: Dima 38 3\nHit `Ctrl+Z` to stop input")
	for {
		stud := readInput()
		studName := stud.Get()
		(*studDB)[studName] = *stud
	}
}

func readInput() *students.Student {
	var (
		name  string
		age   int
		grade int
	)

	_, err := fmt.Scanf("%s %d %d", &name, &age, &grade)
	if err != nil {
		log.Println(err)
	}
	stud := newStudent(name, age, grade)

	return &stud
}

func newStudent(name string, age int, grade int) students.Student {
	stud := students.Student{}
	stud.Put(name, age, grade)
	return stud
}
