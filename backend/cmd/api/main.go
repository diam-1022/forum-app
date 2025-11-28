package main

import (
	"fmt"
	"log"

	"forum-app/backend/internal/config"
	"forum-app/backend/internal/db"
	"forum-app/backend/internal/handlers"
	"forum-app/backend/internal/repository"
	"forum-app/backend/internal/routes"
)

func main() {
	cfg := config.MustLoad()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.MySQLUser, cfg.MySQLPass, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDB, cfg.MySQLParams)
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	repo := repository.New(database)
	handler := handlers.New(repo)
	router := routes.Setup(handler)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
