package service

import "github.com/potterbl/agent/internal/repository"

// Service содержит все сервисы
type Service struct {
	User UserService
	// Здесь можно добавить другие сервисы
	// Product ProductService
	// Order OrderService
}

// NewService создает новый экземпляр Service
func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo.User),
		// Инициализация других сервисов
	}
}
