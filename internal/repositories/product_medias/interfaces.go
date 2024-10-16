package product_medias

import "context"

type Repository interface {
	Creates(ctx context.Context, input CreatesInput) (err error)
}
