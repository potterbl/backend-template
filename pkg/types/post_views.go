package types

import "time"

type PostView struct {
	PostID    int64      `json:"post_id" db:"post_id"`
	UserID    int64      `json:"user_id" db:"user_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
