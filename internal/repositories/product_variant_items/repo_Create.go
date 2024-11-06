package product_variant_items

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/internal/models"
	"time"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	timeNow := time.Now().UTC()
	query := r.sq.Insert("product_variant_items").Columns(
		"product_id", "product_variant_value_1_id", "product_variant_value_2_id",
		"is_primary_product", "price", "stock", "sku",
		"weight", "package_length", "package_width", "package_height", "dimensional_weight",
		"is_active", "image", "created_at", "updated_at",
	).Values(
		input.Data.ProductID, input.Data.ProductVariantValue1ID, input.Data.ProductVariantValue2ID,
		input.Data.IsPrimaryProduct, input.Data.Price, input.Data.Stock, input.Data.Sku,
		input.Data.Weight, input.Data.PackageLength, input.Data.PackageWidth, input.Data.PackageHeight, input.Data.DimensionalWeight,
		input.Data.IsActive, input.Data.Image, timeNow, timeNow,
	).Suffix("RETURNING id")

	rdbms := input.Tx
	if input.Tx == nil {
		rdbms = r.rdbms
	}

	err = rdbms.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeDefault, &output.ID)
	if err != nil {
		return output, collection.Err(err)
	}
	return
}

type CreateInput struct {
	Tx   wsqlx.ReadQuery
	Data models.ProductVariantItem
}

type CreateOutput struct {
	ID int64
}
