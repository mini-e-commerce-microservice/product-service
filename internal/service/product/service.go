package product

import (
	s3wrapper "github.com/SyaibanAhmadRamadhan/go-s3-wrapper"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_medias"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variant_items"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variant_values"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variants"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/products"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/sub_category_items"
)

type service struct {
	subCategoryItemRepository     sub_category_items.Repository
	productRepository             products.Repository
	productMediaRepository        product_medias.Repository
	productVariantRepository      product_variants.Repository
	productVariantItemRepository  product_variant_items.Repository
	productVariantValueRepository product_variant_values.Repository
	s3                            s3wrapper.S3Client
	dbTransaction                 wsqlx.Tx

	// config
	minioConf *secret_proto.Minio
}

type ServiceOption struct {
	SubCategoryItemRepository     sub_category_items.Repository
	ProductRepository             products.Repository
	ProductMediaRepository        product_medias.Repository
	ProductVariantRepository      product_variants.Repository
	ProductVariantItemRepository  product_variant_items.Repository
	ProductVariantValueRepository product_variant_values.Repository
	S3                            s3wrapper.S3Client
	DBTransaction                 wsqlx.Tx

	// config
	MinioConf *secret_proto.Minio
}

func New(opt ServiceOption) *service {
	return &service{
		subCategoryItemRepository:     opt.SubCategoryItemRepository,
		productRepository:             opt.ProductRepository,
		productMediaRepository:        opt.ProductMediaRepository,
		productVariantRepository:      opt.ProductVariantRepository,
		productVariantItemRepository:  opt.ProductVariantItemRepository,
		productVariantValueRepository: opt.ProductVariantValueRepository,
		s3:                            opt.S3,
		dbTransaction:                 opt.DBTransaction,
		minioConf:                     opt.MinioConf,
	}
}
