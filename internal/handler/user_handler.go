package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/potterbl/story-backend/internal/service"
	"github.com/potterbl/story-backend/pkg/types"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// RegisterUserRoutes регистрирует роуты для работы с пользователями
func (h *UserHandler) RegisterUserRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.POST("", h.createUser)
		users.GET("", h.getAllUsers)
		users.GET("/:id", h.getUserByID)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
	}
}

// createUser создает нового пользователя
// @Summary Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body types.CreateUserRequest true "User data"
// @Success 201 {object} types.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [post]
func (h *UserHandler) createUser(c *gin.Context) {
	var req types.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// getAllUsers получает список всех пользователей
// @Summary Get all users
// @Tags users
// @Produce json
// @Success 200 {object} types.UsersListResponse
// @Failure 500 {object} map[string]string
// @Router /api/v1/users [get]
func (h *UserHandler) getAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// getUserByID получает пользователя по ID
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} types.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) getUserByID(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// updateUser обновляет данные пользователя
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body types.UpdateUserRequest true "User data"
// @Success 200 {object} types.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) updateUser(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req types.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(id, &req)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// deleteUser удаляет пользователя
// @Summary Delete user
// @Tags users
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) deleteUser(c *gin.Context) {
	id, err := parseID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// parseID конвертирует строку в uint для ID
func parseID(idStr string) (uint, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, errors.New("invalid ID format")
	}
	return uint(id), nil
}
