package models

import "time"

type ProductVariantItem struct {
	ID                     int64      `db:"id"`
	ProductID              int64      `db:"product_id"`
	ProductVariantValue1ID *int64     `db:"product_variant_value_1_id"`
	ProductVariantValue2ID *int64     `db:"product_variant_value_2_id"`
	IsPrimaryProduct       bool       `db:"is_primary_product"`
	Price                  float64    `db:"price"`
	Stock                  int32      `db:"stock"`
	Sku                    *string    `db:"sku"`
	Weight                 int32      `db:"weight"`
	PackageLength          float64    `db:"package_length"`
	PackageWidth           float64    `db:"package_width"`
	PackageHeight          float64    `db:"package_height"`
	DimensionalWeight      float64    `db:"dimensional_weight"`
	IsActive               bool       `db:"is_active"`
	Image                  *string    `db:"image"`
	CreatedAt              time.Time  `db:"created_at"`
	UpdatedAt              time.Time  `db:"updated_at"`
	DeletedAt              *time.Time `db:"deleted_at"`
}
