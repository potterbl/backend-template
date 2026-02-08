package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/potterbl/story-backend/pkg/types"
)

type UserRepository interface {
	CreateUser(user *types.User) error
	GetUserByID(userID int64) (user types.User, err error)
	DeleteUser(userID int64) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *types.User) error {
	query := `
		INSERT INTO users (
			user_id, 
			username
		) 
		VALUES (
			:user_id, 
			:username
		)`

	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *userRepository) GetUserByID(userID int64) (user types.User, err error) {
	query := `
		SELECT
			user_id,
			username,
			created_at,
			updated_at,
			deleted_at
		FROM users
		WHERE user_id = $1 AND deleted_at IS NULL
	`
	err = r.db.Get(&user, query, userID)
	if err != nil {
		return user, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (r *userRepository) DeleteUser(userID int64) error {
	query := `
		UPDATE users
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE user_id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
