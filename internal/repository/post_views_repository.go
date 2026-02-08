package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostViewsRepository interface {
	AddPostView(postID, userID int64) error
}

type postViewsRepository struct {
	db *sqlx.DB
}

func NewPostViewsRepository(db *sqlx.DB) PostViewsRepository {
	return &postViewsRepository{
		db: db,
	}
}

func (r *postViewsRepository) AddPostView(postID, userID int64) error {
	query := `
		INSERT INTO post_views (
			post_id,
			user_id
		)
		VALUES (
			$post_id,
			$user_id
		)
		ON DUPLICATE DO NOTHING
	`

	_, err := r.db.Exec(query, postID, userID)
	if err != nil {
		return fmt.Errorf("Failed to add post view: %v", err)
	}

	return nil
}
