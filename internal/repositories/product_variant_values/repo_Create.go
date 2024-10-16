package product_variant_values

import (
	"context"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	return
}

type CreateInput struct {
	Data model.ProductVariantValue
}

type CreateOutput struct {
	ID int64
}
