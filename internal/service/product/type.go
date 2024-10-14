package product

import (
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/primitive"
)

type CreateProductInput struct {
	SubCategoryItemID int64
	Condition         string
	VariantName1      null.String
	VariantName2      null.String
	ProductItems      []CreateProductInputProductItem

	// fill by business logic
	isUsedVariant bool
}

type CreateProductInputProductItem struct {
	VariantValue1    null.String
	VariantValue2    null.String
	IsPrimaryProduct bool
	Medias           []CreateProductInputProductItemMedia
}

type CreateProductInputProductItemMedia struct {
	FileUpload primitive.PresignedFileUpload
	IsPrimary  bool
}

type CreateProductOutput struct {
	ID          int64
	MediaUpload []primitive.PresignedFileUploadOutput
}
