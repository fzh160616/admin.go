package main

import (
	"log"
	"os"

	"github.com/fzh160616/admin.go/internal/router"
)

func main() {
	r := router.New()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
