package models

import "time"

type Product struct {
	ID               int64   `db:"id"`
	OutletID         int64   `db:"outlet_id"`
	Name             string  `db:"name"`
	Description      string  `db:"description"`
	ProductCondition string  `db:"product_condition"`
	MinimumPurchase  int32   `db:"minimum_purchase"`
	SizeGuideImage   *string `db:"size_guide_image"`
	IsUsedVariant    bool    `db:"is_used_variant"`

	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
