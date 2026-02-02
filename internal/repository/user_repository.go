package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/potterbl/story-backend/pkg/types"
)

type UserRepository interface {
	CreateUser(user *types.User) error
	GetUserByID(id uint) (*types.User, error)
	GetAllUsers() ([]types.User, error)
	UpdateUser(id uint, user *types.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser создает нового пользователя
func (r *userRepository) CreateUser(user *types.User) error {
	query := `
		INSERT INTO users (email, name, created_at, updated_at) 
		VALUES ($1, $2, NOW(), NOW()) 
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(query, user.Email, user.Name).Scan(
		&user.ID, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// GetUserByID получает пользователя по ID
func (r *userRepository) GetUserByID(id uint) (*types.User, error) {
	user := &types.User{}
	query := `SELECT id, email, name, created_at, updated_at FROM users WHERE id = $1`

	err := r.db.Get(user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// GetAllUsers получает всех пользователей
func (r *userRepository) GetAllUsers() ([]types.User, error) {
	var users []types.User
	query := `SELECT id, email, name, created_at, updated_at FROM users ORDER BY id`

	err := r.db.Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

// UpdateUser обновляет данные пользователя
func (r *userRepository) UpdateUser(id uint, user *types.User) error {
	query := `
		UPDATE users 
		SET email = $1, name = $2, updated_at = NOW() 
		WHERE id = $3 
		RETURNING updated_at`

	err := r.db.QueryRow(query, user.Email, user.Name, id).Scan(&user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to update user: %w", err)
	}

	user.ID = id
	return nil
}

// DeleteUser удаляет пользователя
func (r *userRepository) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
