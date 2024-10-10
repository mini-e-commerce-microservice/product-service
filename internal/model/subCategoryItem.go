package model

import "time"

type SubCategoryItems struct {
	ID            int64      `db:"id"`
	CategoryID    int64      `db:"category_id"`
	SubCategoryID int64      `db:"sub_category_id"`
	Name          string     `db:"name"`
	SizeGuide     bool       `db:"size_guide"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
	DeletedAt     *time.Time `db:"deleted_at"`
}
