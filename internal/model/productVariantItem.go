package model

import "time"

type ProductVariantItem struct {
	ID                int64      `db:"id"`
	ProductID         int64      `db:"product_id"`
	IsPrimaryProduct  bool       `db:"is_primary_product"`
	Price             float64    `db:"price"`
	Stock             int        `db:"stock"`
	Sku               *string    `db:"sku"`
	Weight            int        `db:"weight"`
	PackageLength     float64    `db:"package_length"`
	PackageWidth      float64    `db:"package_width"`
	PackageHeight     float64    `db:"package_height"`
	DimensionalWeight float64    `db:"dimensional_weight"`
	IsActive          bool       `db:"is_active"`
	CreatedAt         time.Time  `db:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at"`
	DeletedAt         *time.Time `db:"deleted_at"`
}

type ProductVariantItemOption struct {
	ProductVariantItemID   int64 `db:"product_variant_item_id"`
	ProductVariantOptionID int64 `db:"product_variant_option_id"`
}
