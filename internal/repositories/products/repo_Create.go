package products

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/internal/models"
	"time"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	timeNow := time.Now().UTC()

	query := r.sq.Insert("products").Columns(
		"name", "outlet_id", "description", "product_condition", "minimum_purchase", "size_guide_image",
		"is_used_variant", "created_at", "updated_at",
	).Values(
		input.Data.Name, input.Data.OutletID, input.Data.Description, input.Data.ProductCondition, input.Data.MinimumPurchase, input.Data.SizeGuideImage,
		input.Data.IsUsedVariant, timeNow, timeNow,
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
	Data models.Product
}

type CreateOutput struct {
	ID int64
}
