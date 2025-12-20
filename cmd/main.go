package main

import (
	"fmt"
	"github.com/potterbl/agent/internal/config"
	"github.com/potterbl/agent/internal/handler"
	"github.com/potterbl/agent/internal/repository"
	"github.com/potterbl/agent/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.InitYAMLConfig()

	_ = repository.NewRepository()
	_ = service.NewService()
	handlers := handler.NewHandler()
	router := handlers.InitRoutes()

	logrus.Printf("Server starting on port %s", cfg.Configuration.Backend.ApiPort)
	if err := router.Run(fmt.Sprintf(":%s", cfg.Configuration.Backend.ApiPort)); err != nil {
		logrus.Fatal("Failed to start server:", err)
	}
}
