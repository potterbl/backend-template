package main

import (
	"fmt"

	"github.com/potterbl/story-backend/internal/config"
	"github.com/potterbl/story-backend/internal/handler"
	"github.com/potterbl/story-backend/internal/repository"
	"github.com/potterbl/story-backend/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.InitYAMLConfig()

	// Инициализация подключения к БД
	db, err := repository.InitDB(cfg.GetDatabaseDSN())
	if err != nil {
		logrus.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Инициализация зависимостей
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	handlers := handler.NewHandler(svc)

	// Инициализация роутов
	router := handlers.InitRoutes()

	logrus.Printf("Server starting on port %s", cfg.Configuration.Backend.ApiPort)
	if err := router.Run(fmt.Sprintf(":%s", cfg.Configuration.Backend.ApiPort)); err != nil {
		logrus.Fatal("Failed to start server:", err)
	}
}
