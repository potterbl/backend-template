package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BoostsRepository interface {
}

type boostsRepository struct {
	db *sqlx.DB
}

func NewBoostsRepository(db *sqlx.DB) BoostsRepository {
	return &boostsRepository{db: db}
}

func (r *boostsRepository) CreateBoost(postID int64, viewsEstimatedCount int) error {
	query := `
		INSERT INTO boosts (
			post_id, 
			views_estimated_count
		)
		VALUES ($1, $2)
	`
	_, err := r.db.Exec(query, postID, viewsEstimatedCount)
	if err != nil {
		return fmt.Errorf("Failed to create boost: %v", err)
	}
	return nil
}
