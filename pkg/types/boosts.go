package types

import "time"

type Boost struct {
	BoostID             int64      `json:"boost_id" db:"boost_id"`
	PostID              int64      `json:"post_id" db:"post_id"`
	ViewsEstimatedCount int64      `json:"views_estimated_count" db:"views_estimated_count"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
