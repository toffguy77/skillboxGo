package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-c:
				{
					fmt.Println("происходит обработка сигнала и выход из программы")
					return
				}
			default:
				{
					fmt.Printf("squared %d: %d\n", i, i*i)
					i++
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	wg.Wait()
}
