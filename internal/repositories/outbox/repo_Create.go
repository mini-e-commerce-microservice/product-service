package outbox

import (
	"context"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (err error) {
	return err
}

type CreateInput struct {
	Data model.Outbox
}
