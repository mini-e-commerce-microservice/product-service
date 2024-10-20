package sub_category_items

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/jmoiron/sqlx"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/primitive"
)

func (r *repository) GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error) {
	query := r.sq.Select("id", "category_id", "sub_category_id", "name", "size_guide").From("sub_category_items").
		OrderByClause("name").OrderByClause("sub_category_id", "DESC")
	queryCount := r.sq.Select("count(*)").From("sub_category_items")

	if input.SubCategoryID.Valid {
		query = query.Where(squirrel.Eq{"sub_category_id": input.SubCategoryID.Int64})
		queryCount = queryCount.Where(squirrel.Eq{"sub_category_id": input.SubCategoryID.Int64})
	}
	if input.CategoryID.Valid {
		query = query.Where(squirrel.Eq{"category_id": input.CategoryID.Int64})
		queryCount = queryCount.Where(squirrel.Eq{"category_id": input.CategoryID.Int64})
	}

	output = GetAllOutput{
		Items: make([]model.SubCategoryItem, 0),
	}
	paginationOutput, err := r.rdbms.QuerySqPagination(ctx, queryCount, query, wsqlx.PaginationInput(input.Pagination), func(rows *sqlx.Rows) (err error) {
		for rows.Next() {
			data := model.SubCategoryItem{}
			err = rows.StructScan(&data)
			if err != nil {
				return collection.Err(err)
			}
			output.Items = append(output.Items, data)
		}

		return
	})
	if err != nil {
		return output, collection.Err(err)
	}

	output.Pagination = primitive.PaginationOutput(paginationOutput)
	return
}

type GetAllInput struct {
	Pagination    primitive.PaginationInput
	SubCategoryID null.Int
	CategoryID    null.Int
}

type GetAllOutput struct {
	Pagination primitive.PaginationOutput
	Items      []model.SubCategoryItem
}
