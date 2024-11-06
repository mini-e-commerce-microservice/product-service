package outlet

import "context"

type Service interface {
	CreateOutlet(ctx context.Context, input CreateOutletInput) (output CreateOutletOutput, err error)
}
