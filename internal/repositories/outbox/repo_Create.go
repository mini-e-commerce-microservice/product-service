package outbox

import (
	"context"
	"encoding/json"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
	"time"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (err error) {
	dataByte, err := json.Marshal(input.Data.Payload)
	if err != nil {
		return collection.Err(err)
	}

	query := r.sq.Insert("outbox").Columns("aggregate_id", "aggregate_type", "payload", "created_at", "trace_parent").
		Values(input.Data.AggregateID, input.Data.AggregateType, string(dataByte), time.Now().UTC(), input.Data.TraceParent)

	rdbms := input.Tx
	if input.Tx == nil {
		rdbms = r.rdbms
	}

	_, err = rdbms.ExecSq(ctx, query)
	if err != nil {
		return collection.Err(err)
	}
	return
}

type CreateInput struct {
	Tx   wsqlx.WriterCommand
	Data model.Outbox
}
