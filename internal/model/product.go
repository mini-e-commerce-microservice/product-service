package model

import "time"

type Products struct {
	ID               int64   `db:"id"`
	Name             string  `db:"name"`
	Description      string  `db:"description"`
	ProductCondition string  `db:"product_condition"`
	MinimumPurchase  int32   `db:"minimum_purchase"`
	SizeGuideImage   *string `db:"size_guide_image"`

	// could be used if no variants
	IsUsedVariant     bool     `db:"is_used_variant"`
	Price             *float64 `db:"price"`
	Stock             *int     `db:"stock"`
	Sku               *string  `db:"sku"`
	Weight            *int     `db:"weight"`
	PackageLength     *float64 `db:"package_length"`
	PackageWidth      *float64 `db:"package_width"`
	PackageHeight     *float64 `db:"package_height"`
	DimensionalWeight *float64 `db:"dimensional_weight"`
	IsActive          *bool    `db:"is_active"`

	TraceParent *string    `db:"trace_parent"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
