package models

type ProductVariantValue struct {
	ID               int64  `db:"id"`
	ProductVariantID int64  `db:"product_variant_id"`
	Value            string `db:"value"`
}
