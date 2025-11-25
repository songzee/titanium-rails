package main

import (
	"log"
	"titanium-rails/internal/config"
	"titanium-rails/internal/server"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	
	log.Printf("Starting server on port %s", cfg.Port)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

