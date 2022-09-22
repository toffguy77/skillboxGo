package main

import (
	"log"
	"net/http"
	"skillbox/middleware"
	"skillbox/pkg/server"
	"sync"
)

func main() {
	s := server.CreateNewServer()
	s.Router.Use(middleware.CommonMiddleware)
	s.MountHandlers()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := http.ListenAndServe(":8080", s.Router)
		if err != nil {
			log.Fatalf("[SERVER] can't start server: %v", err)
		}
	}()

	wg.Wait()
}
