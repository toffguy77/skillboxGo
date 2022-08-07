package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		i := 0
		for {
			fmt.Printf("squared %d: %d\n", i, i*i)
			i++
			time.Sleep(1 * time.Second)
		}
	}()

	s := <-c
	go func() {
		fmt.Printf("происходит обработка сигнала %v и выход из программы\n", s)
		os.Exit(0)
	}()

}
