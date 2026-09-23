package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/kinomfx/student-api/internal/config"
)

func main() {
	fmt.Println("welcome to student api")

	//load config
	cfg := config.MustLoad()

	//database setup

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})
	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	fmt.Printf("Server Started : %s", cfg.HttpServer.Address)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start")
		}
	}()

	wg.Wait()
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("failed to shutdown %s", err.Error())
	}
	fmt.Println("server shutdown successfully")

}
