package product

import (
	"context"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/tracer"
)

// CreateProduct
// List error: ErrOnlyChooseOnePrimaryProduct, ErrOnlyChooseOnePrimaryMedia, ErrMustHavePrimaryMedia, ErrMustHavePrimaryProduct
func (s *service) CreateProduct(ctx context.Context, input CreateProductInput) (output CreateProductOutput, err error) {
	err = s.validateCreateProduct(ctx, &input)
	if err != nil {
		return output, tracer.Error(err)
	}
	return
}

func (s *service) validateCreateProduct(ctx context.Context, input *CreateProductInput) error {
	havePrimaryProduct := false
	primaryProductIsExist := false

	for _, item := range input.ProductItems {
		// Check for primary product
		if item.IsPrimaryProduct {
			if primaryProductIsExist {
				return tracer.Error(ErrOnlyChooseOnePrimaryProduct)
			}
			havePrimaryProduct = true
			primaryProductIsExist = true
		}

		// Check for primary media
		havePrimaryMedia := false
		primaryMediaIsExist := false
		for _, media := range item.Medias {
			if media.IsPrimary {
				if primaryMediaIsExist {
					return tracer.Error(ErrOnlyChooseOnePrimaryMedia)
				}
				havePrimaryMedia = true
				primaryMediaIsExist = true
			}
		}

		// Ensure each product item has a primary media
		if !havePrimaryMedia {
			return tracer.Error(ErrMustHavePrimaryMedia)
		}
	}

	// Ensure there is at least one primary product
	if !havePrimaryProduct {
		return tracer.Error(ErrMustHavePrimaryProduct)
	}

	// validate sub category item and ensure using size guide if sub category item size_guide is true

	// Check if variant is used
	if input.VariantName1.Valid {
		input.isUsedVariant = true
	}

	return nil
}
