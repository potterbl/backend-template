package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	// router := gin.New()
	// router.Use(middleware.Logger())
	// router.Use(middleware.Recovery())
	// router.Use(middleware.CORS())

	// // API группа
	return router
}
