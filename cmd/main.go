package main

import (
	"log"
	http2 "net/http"

	"tinyURL/internal/config"
	"tinyURL/internal/http"
	"tinyURL/internal/shortner"
	"tinyURL/internal/storage/memory"
)

func main() {
	cfg := config.Load()

	//data storage
	store := memory.New()

	//svc
	svc := shortner.New(store)
	router := http.NewRouter(svc)

	log.Printf("Starting server on %s", cfg.Server.Address)
	if err := http2.ListenAndServe(cfg.Server.Address, router); err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}

}
