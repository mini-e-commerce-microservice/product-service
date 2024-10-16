package sub_category_items

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/tracer"
)

func (r *repository) Get(ctx context.Context, input GetInput) (output GetOutput, err error) {
	query := r.sq.Select("id", "category_id", "sub_category_id", "name", "size_guide").From("sub_category_items")
	if input.ID.Valid {
		query = query.Where(squirrel.Eq{"id": input.ID.Int64})
	}

	err = r.rdbms.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeStruct, &output.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repositories.ErrDataNotFound
		}
		return output, tracer.Error(err)
	}
	return
}

type GetInput struct {
	ID null.Int
}

type GetOutput struct {
	Data model.SubCategoryItem
}
