package main

import (
	"context"
	"ferry-learn/database"
	"ferry-learn/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	database.Setup()
	server := &http.Server{
		Addr:    ":11451",
		Handler: router.InitRouter(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("listen: ", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("server shuting down")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("server shutdown: ", err)
	}
}
