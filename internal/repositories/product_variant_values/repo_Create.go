package product_variant_values

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	query := r.sq.Insert("product_variant_values").Columns(
		"product_variant_id", "value",
	).Values(
		input.Data.ProductVariantID, input.Data.Value,
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
	Data model.ProductVariantValue
}

type CreateOutput struct {
	ID int64
}
