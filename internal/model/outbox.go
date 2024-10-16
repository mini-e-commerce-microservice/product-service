package model

import "time"

type AggregateType string

const AggregateTypeProduct AggregateType = "aggregate_product"

type Outbox struct {
	ID            int64         `db:"id"`
	AggregateID   int64         `db:"aggregate_id"`
	AggregateType AggregateType `db:"aggregate_type"`
	Payload       any           `db:"payload"`
	TraceParent   *string       `db:"trace_parent"`
	CreatedAt     time.Time     `db:"created_at"`
}

type OutboxPayloadProduct struct {
	ID                  int64                       `json:"id"`
	Variant1            OutboxPayloadProductVariant `json:"variant_1"`
	Variant2            OutboxPayloadProductVariant `json:"variant_2"`
	SubCategoryItemName string                      `json:"sub_category_item_name"`
	Name                string                      `json:"name"`
	Description         string                      `json:"description"`
	Price               float64                     `json:"price"`
	Stock               int32                       `json:"stock"`
	Sku                 *string                     `json:"sku"`
	Weight              int64                       `json:"weight"`
	PackageLength       float64                     `json:"package_length"`
	PackageWidth        float64                     `json:"package_width"`
	PackageHeight       float64                     `json:"package_height"`
	DimensionalWeight   float64                     `json:"dimensional_weight"`
	IsActive            bool                        `json:"is_active"`
	ProductCondition    string                      `json:"product_condition"`
	MinimumPurchase     int32                       `json:"minimum_purchase"`
	SizeGuideImage      *string                     `json:"size_guide_image"`
	IsUsedVariant       bool                        `json:"is_used_variant"`
	CreatedAt           time.Time                   `json:"created_at"`
	UpdatedAt           time.Time                   `json:"updated_at"`
	DeletedAt           *time.Time                  `json:"deleted_at"`
}

type OutboxPayloadProductVariant struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
