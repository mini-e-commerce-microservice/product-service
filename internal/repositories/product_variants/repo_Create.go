package product_variants

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/internal/models"
	"time"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	timeNow := time.Now().UTC()

	query := r.sq.Insert("product_variants").Columns(
		"product_id", "name", "created_at", "updated_at",
	).Values(
		input.Data.ProductID, input.Data.Name, timeNow, timeNow,
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
	Data models.ProductVariant
}

type CreateOutput struct {
	ID int64
}
