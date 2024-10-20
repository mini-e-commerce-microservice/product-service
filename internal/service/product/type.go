package product

import (
	"github.com/SyaibanAhmadRamadhan/go-collection/generic"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/internal/model"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/primitive"
)

type CreateProductInput struct {
	UserID            int64
	SubCategoryItemID int64
	Condition         string
	MinimumPurchase   int32
	Description       string
	Name              string
	VariantName1      null.String
	VariantName2      null.String
	ProductItems      []CreateProductInputProductItem
	SizeGuide         null.Value[primitive.PresignedFileUpload]
	Medias            []CreateProductInputProductMedia

	// fill by business logic
	isUsedVariant      bool
	sizeGuideImageName *string
	subCategoryItem    model.SubCategoryItem
}

type CreateProductInputProductItem struct {
	VariantValue1    null.String
	VariantValue2    null.String
	Price            float64
	Stock            int32
	SKU              null.String
	Weight           int32
	PackageLength    float64
	PackageWidth     float64
	PackageHeight    float64
	IsPrimaryProduct bool
	IsActive         bool
	Image            null.Value[primitive.PresignedFileUpload]

	// fill by business logic
	dimensionalWeight float64
}

func (item *CreateProductInputProductItem) calculateDimensionalWeight(volumetricFactor float64) {
	if volumetricFactor == 0 {
		volumetricFactor = 5000.0
	}
	item.dimensionalWeight = (item.PackageLength * item.PackageWidth * item.PackageHeight) / volumetricFactor
}

type CreateProductInputProductMedia struct {
	FileUpload primitive.PresignedFileUpload
	IsPrimary  bool
}

type CreateProductOutput struct {
	ID                   int64
	OptionalImageUploads []null.Value[primitive.PresignedFileUploadOutput]
	MediaUploads         []primitive.PresignedFileUploadOutput
}

type createPresignedUploadMediaProductInput struct {
	fileUploads    []primitive.PresignedFileUpload
	optionalImages []null.Value[primitive.PresignedFileUpload]
}

type createPresignedUploadMediaProductOutput struct {
	mediaUploads         []primitive.PresignedFileUploadOutput
	optionalImageUploads []null.Value[primitive.PresignedFileUploadOutput]
}

type insertProductMediaInput struct {
	tx        wsqlx.WriterCommand
	productID int64
	items     []insertProductMediaInputItem
}

type insertProductMediaInputItem struct {
	fileUpload primitive.PresignedFileUpload
	isPrimary  bool
}

type insertProductVariantValuesInput struct {
	tx                   wsqlx.Rdbms
	productVariantID1    int64
	productVariantID2    int64
	productVariantValue1 []string
	productVariantValue2 []string
}

type insertProductVariantValuesOutput struct {
	// key is value variant, value is value variant
	productVariantValues *generic.SafeMap[string, int64]
}
