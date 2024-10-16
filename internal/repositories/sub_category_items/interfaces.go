package sub_category_items

import "context"

type Repository interface {
	Get(ctx context.Context, input GetInput) (output GetOutput, err error)
	GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error)
}
