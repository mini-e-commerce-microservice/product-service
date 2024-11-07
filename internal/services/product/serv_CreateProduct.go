package product

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"github.com/SyaibanAhmadRamadhan/go-collection/generic"
	s3wrapper "github.com/SyaibanAhmadRamadhan/go-s3-wrapper"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/internal/models"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/outbox"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/outlets"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_medias"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variant_items"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variant_values"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variants"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/products"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/sub_category_items"
	"github.com/mini-e-commerce-microservice/product-service/internal/util"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/primitive"
	"golang.org/x/sync/errgroup"
	"time"
)

// CreateProduct
//
// List error: ErrOnlyChooseOnePrimaryProduct, ErrOnlyChooseOnePrimaryMedia, ErrMustHavePrimaryMedia, ErrMustHavePrimaryProduct,
// ErrInvalidSubCategoryItem, ErrMustHaveSizeGuide, ErrVariantValue1IsRequired, ErrVariantValue2IsRequired, ErrOutletNotFound
func (s *service) CreateProduct(ctx context.Context, input CreateProductInput) (output CreateProductOutput, err error) {
	err = s.validateCreateProduct(ctx, &input)
	if err != nil {
		return output, collection.Err(err)
	}

	outletOutput, err := s.outletRepository.FindOne(ctx, outlets.FindOneInput{
		UserID: null.IntFrom(input.UserID),
	})
	if err != nil {
		if errors.Is(err, repositories.ErrDataNotFound) {
			err = errors.Join(err, ErrOutletNotFound)
		}
		return output, collection.Err(err)
	}

	createPresignedUploadOutput, err := s.createPresignedUploadMediaProduct(ctx, input.toCreatePresignedUploadMediaProductInput())
	if err != nil {
		return output, collection.Err(err)
	}
	output.MediaUploads = createPresignedUploadOutput.mediaUploads
	output.OptionalImageUploads = createPresignedUploadOutput.optionalImageUploads

	err = s.dbTransaction.DoTxContext(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted},
		func(ctx context.Context, tx wsqlx.Rdbms) (err error) {
			outputCreateProduct, err := s.productRepository.Create(ctx, products.CreateInput{
				Tx: tx,
				Data: models.Product{
					OutletID:         outletOutput.Data.ID,
					Name:             input.Name,
					Description:      input.Description,
					ProductCondition: input.Condition,
					MinimumPurchase:  input.MinimumPurchase,
					SizeGuideImage:   input.sizeGuideImageName,
					IsUsedVariant:    input.isUsedVariant,
				},
			})
			if err != nil {
				return collection.Err(err)
			}
			output.ID = outputCreateProduct.ID

			var eg errgroup.Group

			eg.Go(func() (err error) {
				err = s.insertProductMedia(ctx, insertProductMediaInput{
					tx:        tx,
					productID: outputCreateProduct.ID,
					items:     input.toInsertProductMediaInputItems(),
				})
				if err != nil {
					return collection.Err(err)
				}
				return
			})

			createProductVariant1Output := product_variants.CreateOutput{}
			if input.VariantName1.Valid {
				eg.Go(func() (err error) {
					createProductVariant1Output, err = s.productVariantRepository.Create(ctx, product_variants.CreateInput{
						Tx: tx,
						Data: models.ProductVariant{
							ProductID: outputCreateProduct.ID,
							Name:      input.VariantName1.String,
						},
					})
					if err != nil {
						return collection.Err(err)
					}
					return
				})
			}
			createProductVariant2Output := product_variants.CreateOutput{}
			if input.VariantName2.Valid {
				eg.Go(func() (err error) {
					createProductVariant2Output, err = s.productVariantRepository.Create(ctx, product_variants.CreateInput{
						Tx: tx,
						Data: models.ProductVariant{
							ProductID: outputCreateProduct.ID,
							Name:      input.VariantName2.String,
						},
					})
					if err != nil {
						return collection.Err(err)
					}
					return
				})
			}

			if err = eg.Wait(); err != nil {
				return collection.Err(err)
			}

			insertVariantValuesInput, err := input.toInsertProductVariantValuesInput(tx, createProductVariant1Output.ID, createProductVariant2Output.ID)
			if err != nil {
				return collection.Err(err)
			}
			insertVariantValuesOutput, err := s.insertProductVariantValues(ctx, insertVariantValuesInput)
			if err != nil {
				return collection.Err(err)
			}

			for _, item := range input.ProductItems {
				eg.Go(func() (err error) {
					item.calculateDimensionalWeight(0)
					var (
						productVariantValue1ID       *int64
						productVariantValue2ID       *int64
						image                        *string
						outboxPayloadProductVariant1 *models.OutboxPayloadProductVariant
						outboxPayloadProductVariant2 *models.OutboxPayloadProductVariant
					)

					val1, ok := insertVariantValuesOutput.productVariantValues.Load(item.VariantValue1.String)
					if ok {
						outboxPayloadProductVariant1 = &models.OutboxPayloadProductVariant{
							Name:  input.VariantName1.String,
							Value: item.VariantValue1.String,
						}
						productVariantValue1ID = &val1
					}
					val2, ok := insertVariantValuesOutput.productVariantValues.Load(item.VariantValue2.String)
					if ok {
						outboxPayloadProductVariant2 = &models.OutboxPayloadProductVariant{
							Name:  input.VariantName2.String,
							Value: item.VariantValue2.String,
						}
						productVariantValue2ID = &val2
					}
					if item.Image.Valid {
						image = &item.Image.V.GeneratedFileName
					}

					productVariantItemCreateOutput, err := s.productVariantItemRepository.Create(ctx, product_variant_items.CreateInput{
						Tx: tx,
						Data: models.ProductVariantItem{
							ProductID:              outputCreateProduct.ID,
							ProductVariantValue1ID: productVariantValue1ID,
							ProductVariantValue2ID: productVariantValue2ID,
							IsPrimaryProduct:       item.IsPrimaryProduct,
							Price:                  item.Price,
							Stock:                  item.Stock,
							Sku:                    item.SKU.Ptr(),
							Weight:                 item.Weight,
							PackageLength:          item.PackageLength,
							PackageWidth:           item.PackageWidth,
							PackageHeight:          item.PackageHeight,
							DimensionalWeight:      item.dimensionalWeight,
							IsActive:               item.IsActive,
							Image:                  image,
						},
					})
					if err != nil {
						return collection.Err(err)
					}

					err = s.outboxRepository.Create(ctx, outbox.CreateInput{
						Tx: tx,
						Data: models.Outbox{
							AggregateID:   productVariantItemCreateOutput.ID,
							AggregateType: string(outbox.AggregateTypeProduct),
							Payload: models.OutboxPayloadProduct{
								ID:                  productVariantItemCreateOutput.ID,
								OutletID:            outletOutput.Data.ID,
								UserID:              input.UserID,
								Variant1:            null.ValueFromPtr(outboxPayloadProductVariant1),
								Variant2:            null.ValueFromPtr(outboxPayloadProductVariant2),
								SubCategoryItemName: input.subCategoryItem.Name,
								Name:                input.Name,
								Description:         input.Name,
								Price:               item.Price,
								Stock:               item.Stock,
								Sku:                 item.SKU.Ptr(),
								Weight:              item.Weight,
								PackageLength:       item.PackageLength,
								PackageWidth:        item.PackageWidth,
								PackageHeight:       item.PackageHeight,
								DimensionalWeight:   item.dimensionalWeight,
								IsActive:            item.IsActive,
								ProductCondition:    input.Condition,
								MinimumPurchase:     input.MinimumPurchase,
								SizeGuideImage:      input.sizeGuideImageName,
								CreatedAt:           time.Now().UTC(),
								UpdatedAt:           time.Now().UTC(),
							},
							TraceParent: util.GetTraceParent(ctx),
						},
					})
					if err != nil {
						return collection.Err(err)
					}
					return
				})
			}

			if err = eg.Wait(); err != nil {
				return collection.Err(err)
			}
			return
		},
	)
	if err != nil {
		return output, collection.Err(err)
	}

	return
}

func (s *service) insertProductVariantValues(ctx context.Context, input insertProductVariantValuesInput) (output insertProductVariantValuesOutput, err error) {
	output = insertProductVariantValuesOutput{
		productVariantValues: &generic.SafeMap[string, int64]{},
	}

	var eg errgroup.Group

	if input.productVariantValue1 != nil && input.productVariantID1 != 0 {
		for _, value1 := range input.productVariantValue1 {
			eg.Go(func() (err error) {
				productVariantValueOutput, err := s.productVariantValueRepository.Create(ctx, product_variant_values.CreateInput{
					Tx: input.tx,
					Data: models.ProductVariantValue{
						ProductVariantID: input.productVariantID1,
						Value:            value1,
					},
				})
				if err != nil {
					return collection.Err(err)
				}

				output.productVariantValues.Store(value1, productVariantValueOutput.ID)
				return
			})
		}
	}

	if input.productVariantValue2 != nil && input.productVariantID2 != 0 {
		for _, value2 := range input.productVariantValue2 {
			eg.Go(func() (err error) {
				productVariantValueOutput, err := s.productVariantValueRepository.Create(ctx, product_variant_values.CreateInput{
					Tx: input.tx,
					Data: models.ProductVariantValue{
						ProductVariantID: input.productVariantID2,
						Value:            value2,
					},
				})
				if err != nil {
					return collection.Err(err)
				}
				output.productVariantValues.Store(value2, productVariantValueOutput.ID)
				return
			})
		}
	}

	if err = eg.Wait(); err != nil {
		return output, collection.Err(err)
	}
	return
}

func (s *service) insertProductMedia(ctx context.Context, input insertProductMediaInput) (err error) {
	productMedias := make([]models.ProductMedia, 0, len(input.items))
	for _, item := range input.items {
		productMedias = append(productMedias, models.ProductMedia{
			ProductID:      input.productID,
			Media:          item.fileUpload.GeneratedFileName,
			MediaType:      item.fileUpload.MimeType.MediaType(),
			IsPrimaryMedia: item.isPrimary,
		})
	}

	err = s.productMediaRepository.Creates(ctx, product_medias.CreatesInput{
		Tx:   input.tx,
		Data: productMedias,
	})
	if err != nil {
		err = collection.Err(err)
	}

	return
}

func (s *service) createPresignedUploadMediaProduct(ctx context.Context, input createPresignedUploadMediaProductInput) (output createPresignedUploadMediaProductOutput, err error) {
	output = createPresignedUploadMediaProductOutput{
		mediaUploads:         make([]primitive.PresignedFileUploadOutput, 0, len(input.fileUploads)),
		optionalImageUploads: make([]null.Value[primitive.PresignedFileUploadOutput], 0, len(input.optionalImages)),
	}

	var eg errgroup.Group

	for _, image := range input.optionalImages {
		if image.Valid {
			eg.Go(func() (err error) {
				outputPresignedUrlUploadObject, err := s.s3.PresignedUrlUploadObject(ctx, s3wrapper.PresignedUrlUploadObjectInput{
					BucketName: s.minioConf.PrivateBucket,
					Path:       image.V.GeneratedFileName,
					MimeType:   string(image.V.MimeType),
					Checksum:   image.V.ChecksumSHA256,
					Expired:    10 * time.Minute,
				})
				if err != nil {
					return collection.Err(err)
				}

				output.optionalImageUploads = append(output.optionalImageUploads, null.ValueFrom(primitive.PresignedFileUploadOutput{
					Identifier:      image.V.Identifier,
					UploadURL:       outputPresignedUrlUploadObject.URL,
					UploadExpiredAt: outputPresignedUrlUploadObject.ExpiredAt,
					MinioFormData:   outputPresignedUrlUploadObject.MinioFormData,
				}))
				return
			})
		}
	}

	for _, media := range input.fileUploads {
		eg.Go(func() (err error) {
			outputPresignedUrlUploadObject, err := s.s3.PresignedUrlUploadObject(ctx, s3wrapper.PresignedUrlUploadObjectInput{
				BucketName: s.minioConf.PrivateBucket,
				Path:       media.GeneratedFileName,
				MimeType:   string(media.MimeType),
				Checksum:   media.ChecksumSHA256,
				Expired:    10 * time.Minute,
			})
			if err != nil {
				return collection.Err(err)
			}
			output.mediaUploads = append(output.mediaUploads, primitive.PresignedFileUploadOutput{
				Identifier:      media.Identifier,
				UploadURL:       outputPresignedUrlUploadObject.URL,
				UploadExpiredAt: outputPresignedUrlUploadObject.ExpiredAt,
				MinioFormData:   outputPresignedUrlUploadObject.MinioFormData,
			})
			return
		})
	}

	if err = eg.Wait(); err != nil {
		return output, collection.Err(err)
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
				return collection.Err(ErrOnlyChooseOnePrimaryProduct)
			}
			havePrimaryProduct = true
			primaryProductIsExist = true
		}

		if input.VariantName1.Valid && !item.VariantValue1.Valid {
			return collection.Err(ErrVariantValue1IsRequired)
		}
		if input.VariantName2.Valid && !item.VariantValue2.Valid {
			return collection.Err(ErrVariantValue2IsRequired)
		}
	}
	// Ensure there is at least one primary product
	if !havePrimaryProduct {
		return collection.Err(ErrMustHavePrimaryProduct)
	}

	// Check for primary media
	havePrimaryMedia := false
	primaryMediaIsExist := false
	for _, media := range input.Medias {
		if media.IsPrimary {
			if primaryMediaIsExist {
				return collection.Err(ErrOnlyChooseOnePrimaryMedia)
			}
			havePrimaryMedia = true
			primaryMediaIsExist = true
		}
	}
	// Ensure each product item has a primary media
	if !havePrimaryMedia {
		return collection.Err(ErrMustHavePrimaryMedia)
	}

	// validate sub category item and ensure using size guide if sub category item size_guide is true
	subCategoryItem, err := s.subCategoryItemRepository.Get(ctx, sub_category_items.GetInput{
		ID: null.IntFrom(input.SubCategoryItemID),
	})
	if err != nil {
		if errors.Is(err, repositories.ErrDataNotFound) {
			err = errors.Join(err, ErrInvalidSubCategoryItem)
		}
		return collection.Err(err)
	}
	if subCategoryItem.Data.SizeGuide && !input.SizeGuide.Valid {
		return collection.Err(ErrMustHaveSizeGuide)
	}
	if subCategoryItem.Data.SizeGuide {
		input.sizeGuideImageName = &input.SizeGuide.V.GeneratedFileName
	}
	input.subCategoryItem = subCategoryItem.Data

	// Check if variant is used
	if input.VariantName1.Valid {
		input.isUsedVariant = true
	}

	return nil
}

func (c CreateProductInput) toInsertProductMediaInputItems() []insertProductMediaInputItem {
	i := make([]insertProductMediaInputItem, 0, len(c.Medias))
	for _, media := range c.Medias {
		i = append(i, insertProductMediaInputItem{
			fileUpload: media.FileUpload,
			isPrimary:  media.IsPrimary,
		})
	}

	return i
}

func (c CreateProductInput) toCreatePresignedUploadMediaProductInput() createPresignedUploadMediaProductInput {
	i := createPresignedUploadMediaProductInput{
		fileUploads: make([]primitive.PresignedFileUpload, 0, len(c.Medias)),
		optionalImages: []null.Value[primitive.PresignedFileUpload]{
			c.SizeGuide,
		},
	}

	for _, item := range c.ProductItems {
		i.optionalImages = append(i.optionalImages, item.Image)
	}
	for _, media := range c.Medias {
		i.fileUploads = append(i.fileUploads, media.FileUpload)
	}

	return i
}

func (c CreateProductInput) toInsertProductVariantValuesInput(tx wsqlx.Rdbms, variantID1, variantID2 int64) (insertProductVariantValuesInput, error) {
	output := insertProductVariantValuesInput{
		tx:                   tx,
		productVariantID1:    variantID1,
		productVariantID2:    variantID2,
		productVariantValue1: make([]string, 0),
		productVariantValue2: make([]string, 0),
	}

	appendUnique := func(slice []string, value string) []string {
		for _, v := range slice {
			if v == value {
				return slice
			}
		}
		return append(slice, value)
	}

	for _, item := range c.ProductItems {
		if variantID1 != 0 {
			if !item.VariantValue1.Valid {
				return insertProductVariantValuesInput{}, collection.Err(ErrVariantValue1IsRequired)
			}
			output.productVariantValue1 = appendUnique(output.productVariantValue1, item.VariantValue1.String)
		}
		if variantID2 != 0 {
			if !item.VariantValue2.Valid {
				return insertProductVariantValuesInput{}, collection.Err(ErrVariantValue2IsRequired)
			}
			output.productVariantValue2 = appendUnique(output.productVariantValue2, item.VariantValue2.String)
		}
	}

	return output, nil
}
