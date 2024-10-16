package product_variant_items

import (
	"context"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	return
}

type CreateInput struct {
	Tx   wsqlx.ReadQuery
	Data model.ProductVariantItem
}

type CreateOutput struct {
	ID int64
}
