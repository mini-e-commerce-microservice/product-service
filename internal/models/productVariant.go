package models

import "time"

type ProductVariant struct {
	ID        int64      `db:"id"`
	ProductID int64      `db:"product_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
