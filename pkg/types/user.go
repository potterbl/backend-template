package types

import "time"

// User представляет пользователя в системе
type User struct {
	ID        uint      `json:"id" db:"id"`
	Email     string    `json:"email" db:"email" binding:"required,email"`
	Name      string    `json:"name" db:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUserRequest представляет запрос на создание пользователя
type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
	Name  string `json:"name" binding:"required"`
}

// UpdateUserRequest представляет запрос на обновление пользователя
type UpdateUserRequest struct {
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
	Name  *string `json:"name,omitempty"`
}

// UserResponse представляет ответ с данными пользователя
type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UsersListResponse представляет ответ со списком пользователей
type UsersListResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
}