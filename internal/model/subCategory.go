package model

import "time"

type SubCategory struct {
	ID         int64      `db:"id"`
	CategoryID int64      `db:"category_id"`
	Name       string     `db:"name"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	DeletedAt  *time.Time `db:"deleted_at"`
}
