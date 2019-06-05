package main

import (
	"go-learn/projects/simple/config"
	"go-learn/projects/simple/model"
	"go-learn/projects/simple/web"

	"log"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to reading config file, %s\n", err)
	}

	service, err := model.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize model for operating all service, %s\n", err)
	}

	server := web.NewServer(cfg, service)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen for http server, %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	log.Println("simple is running")
	<-quit
	log.Println("simple is stopped")
}
