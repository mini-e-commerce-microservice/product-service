package outlets

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/internal/models"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (output CreateOutput, err error) {
	columns, values := collection.GetTagsWithValues(input.Data, "db", "id")
	query := r.sq.Insert("outlets").Columns(columns...).Values(values...).Suffix("RETURNING id")

	rdbms := input.Tx
	if input.Tx == nil {
		rdbms = r.rdbms
	}

	err = rdbms.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeDefault, &output.ID)
	if err != nil {
		return output, collection.Err(err)
	}
	return
}

type CreateInput struct {
	Tx   wsqlx.ReadQuery
	Data models.Outlet
}
type CreateOutput struct {
	ID int64
}
