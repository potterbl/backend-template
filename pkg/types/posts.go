package types

import "time"

type Post struct {
	PostID      int64      `json:"post_id" db:"post_id"`
	UserID      int64      `json:"user_id" db:"user_id"`
	ImageName   string     `json:"image_name" db:"image_name"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
