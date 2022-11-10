package main

import (
	"context"
	"log"
	"net/http"
	"skillbox/internal/flags"
	"skillbox/internal/server"
	"skillbox/middleware"
	"sync"
)

func main() {
	ctx := context.Background()
	var (
		dataCtx  flags.DataType = "dataType"
		userData                = flags.Data{}
	)
	ctx = context.WithValue(ctx, dataCtx, &userData)
	err := flags.ParseUserFlags(&ctx)
	if err != nil {
		log.Fatalf("can't parse user flags: %v", err)
	}

	log.Println("starting http server...")
	s := server.CreateNewServer()
	s.Router.Use(middleware.CommonMiddleware)
	s.MountHandlers()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("ready to serve")
		err := http.ListenAndServe(":"+userData.PORT1, s.Router)
		if err != nil {
			log.Fatalf("[SERVER] can't start server: %v", err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("ready to serve")
		err := http.ListenAndServe(":"+userData.PORT2, s.Router)
		if err != nil {
			log.Fatalf("[SERVER] can't start server: %v", err)
		}
	}()

	wg.Wait()
}
