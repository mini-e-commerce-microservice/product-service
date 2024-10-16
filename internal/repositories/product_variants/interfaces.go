package product_variants

import "context"

type Repository interface {
	Create(ctx context.Context, input CreateInput) (output CreateOutput, err error)
}
