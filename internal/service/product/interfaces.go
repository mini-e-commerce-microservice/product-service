package product

import "context"

type Service interface {
	CreateProduct(ctx context.Context, input CreateProductInput) (output CreateProductOutput, err error)
}
