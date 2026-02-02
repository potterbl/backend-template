package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/potterbl/story-backend/internal/service"
)

// Handler содержит все хендлеры
type Handler struct {
	User *UserHandler
	// Здесь можно добавить другие хендлеры
	// Product *ProductHandler
	// Order *OrderHandler
}

// NewHandler создает новый экземпляр Handler
func NewHandler(service *service.Service) *Handler {
	return &Handler{
		User: NewUserHandler(service.User),
		// Инициализация других хендлеров
	}
}

// InitRoutes инициализирует все роуты
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Health check endpoint
	router.GET("/health", h.healthCheck)

	// API группа
	api := router.Group("/api/v1")
	{
		// Регистрация роутов для пользователей
		h.User.RegisterUserRoutes(api)

		// Здесь можно добавить другие роуты
		// h.Product.RegisterProductRoutes(api)
		// h.Order.RegisterOrderRoutes(api)
	}

	return router
}

// healthCheck проверка здоровья сервиса
func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "backend-template",
	})
}
