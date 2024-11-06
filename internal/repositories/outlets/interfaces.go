package outlets

import "context"

type Repository interface {
	Create(ctx context.Context, input CreateInput) (output CreateOutput, err error)
	FindOne(ctx context.Context, input FindOneInput) (output FindOneOutput, err error)
}
