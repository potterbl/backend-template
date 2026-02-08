package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/potterbl/story-backend/pkg/types"
)

const (
	boostPostsPercentagePerFeed   = 60
	defaultPostsPercentagePerFeed = 40
)

type PostsRepository interface {
	CreatePost(post *types.Post) error
	GetPostsByUserID(userID int64) (posts []types.Post, err error)
	GetPostByID(postID int64) (post types.Post, err error)
	DeletePost(postID int64) error
}

type postsRepository struct {
	db *sqlx.DB
}

func NewPostsRepository(db *sqlx.DB) PostsRepository {
	return &postsRepository{
		db: db,
	}
}

func (r *postsRepository) CreatePost(post *types.Post) error {
	query := `
		INSERT INTO posts (
			user_id,
			image_name,
			description
		)
		VALUES (
			:user_id,
			:image_name,
			:description
		)
	`

	_, err := r.db.NamedExec(query, post)
	if err != nil {
		return fmt.Errorf("Failed to create post: %v", err)
	}

	return nil
}

func (r *postsRepository) GetPostsByUserID(userID int64) (posts []types.Post, err error) {
	query := `
		SELECT
			post_id,
			user_id,
			image_name,
			description,
			created_at,
			updated_at
		FROM posts
		WHERE user_id = $1 AND deleted_at IS NULL
	`
	err = r.db.Select(&posts, query, userID)
	if err != nil {
		return posts, fmt.Errorf("Failed to get posts by user ID: %v", err)
	}

	return posts, nil
}

func (r *postsRepository) GetPostByID(postID int64) (post types.Post, err error) {
	query := `
		SELECT
			post_id,
			user_id,
			image_name,
			description,
			created_at,
			updated_at
		FROM posts
		WHERE post_id = $1 AND deleted_at IS NULL
	`
	err = r.db.Get(&post, query, postID)
	if err != nil {
		return post, fmt.Errorf("Failed to get post by ID: %v", err)
	}

	return post, nil
}

func (r *postsRepository) DeletePost(postID int64) error {
	query := `
		UPDATE posts
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE post_id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, postID)
	if err != nil {
		return fmt.Errorf("Failed to delete post: %v", err)
	}

	return nil
}

func (r *postsRepository) GetFeedPosts(userID int64, limit int64) (posts []types.Post, err error) {
	var boostPosts []types.Post
	var defaultPosts []types.Post

	boostPostsQuery := `
		SELECT
			p.post_id,
			p.user_id,
			p.image_name,
			p.description,
			p.created_at,
			p.updated_at
		FROM posts p
		LEFT JOIN post_views pv ON pv.post_id = p.post_id
			AND pv.user_id = $1
		JOIN boosts b ON b.post_id = p.post_id
			AND b.deleted_at IS NULL
		WHERE 
			pv.post_id IS NULL
			AND p.deleted_at IS NULL
			AND pv.user_id IS NULL
		ORDER BY p.created_at DESC
		LIMIT $2
	`
	err = r.db.Select(&boostPosts, boostPostsQuery, userID, (limit/100)*boostPostsPercentagePerFeed)
	if err != nil {
		return posts, fmt.Errorf("Failed to get feed posts: %v", err)
	}

	defaultPostsQuery := `
		SELECT
			p.post_id,
			p.user_id,
			p.image_name,
			p.description,
			p.created_at,
			p.updated_at
		FROM posts p
		LEFT JOIN post_views pv ON pv.post_id = p.post_id
			AND pv.user_id = $1
		LEFT JOIN boosts b ON b.post_id = p.post_id
			AND b.deleted_at IS NULL
		WHERE 
			p.deleted_at IS NULL
			AND pv.user_id IS NULL
			AND b.post_id IS NULL
		ORDER BY p.created_at DESC
		LIMIT $2
	`
	err = r.db.Select(&defaultPosts, defaultPostsQuery, userID, (limit/100)*defaultPostsPercentagePerFeed)
	if err != nil {
		return posts, fmt.Errorf("Failed to get feed posts: %v", err)
	}

	if int64(len(boostPosts)) < (limit/100)*boostPostsPercentagePerFeed {
		remainingBoostPostsLimit := (limit/100)*boostPostsPercentagePerFeed - int64(len(boostPosts))
		err = r.db.Select(&boostPosts, defaultPostsQuery, userID, remainingBoostPostsLimit)
		if err != nil {
			return posts, fmt.Errorf("Failed to get feed posts: %v", err)
		}
	}

	posts = append(posts, boostPosts...)
	posts = append(posts, defaultPosts...)

	return posts, nil
}
