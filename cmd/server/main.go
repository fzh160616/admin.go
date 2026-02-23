package main

import (
	"log"

	"github.com/fzh160616/admin.go/internal/config"
	"github.com/fzh160616/admin.go/internal/router"
	"github.com/fzh160616/admin.go/internal/store"
)

func main() {
	cfg := config.Load()

	db, err := store.NewMySQL(cfg.MySQLDSN)
	if err != nil {
		log.Fatalf("db init failed: %v", err)
	}

	r := router.New(db, cfg.JWTSecret)

	addr := ":" + cfg.AppPort
	log.Printf("server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
