package outlets

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/internal/models"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories"
)

func (r *repository) FindOne(ctx context.Context, input FindOneInput) (output FindOneOutput, err error) {
	query := r.sq.Select("*").From("outlets")
	if input.ID.Valid {
		query = query.Where(squirrel.Eq{"id": input.ID.Int64})
	}
	if input.UserID.Valid {
		query = query.Where(squirrel.Eq{"user_id": input.UserID.Int64})
	}

	err = r.rdbms.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeStruct, &output.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = repositories.ErrDataNotFound
		}
		return output, collection.Err(err)
	}

	return
}

type FindOneInput struct {
	ID     null.Int
	UserID null.Int
}

type FindOneOutput struct {
	Data models.Outlet
}
