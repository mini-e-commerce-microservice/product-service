package outbox

import "context"

type Repository interface {
	Create(ctx context.Context, input CreateInput) (err error)
}
