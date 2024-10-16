package product_variant_values

import "context"

type Repository interface {
	Create(ctx context.Context, input CreateInput) (output CreateOutput, err error)
}
