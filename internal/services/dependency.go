package services

import (
	"context"
	s3_wrapper_minio "github.com/SyaibanAhmadRamadhan/go-s3-wrapper/minio"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/product-service/internal/conf"
	"github.com/mini-e-commerce-microservice/product-service/internal/infra"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/outbox"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/outlets"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_medias"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variant_items"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variant_values"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/product_variants"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/products"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/sub_category_items"
	"github.com/mini-e-commerce-microservice/product-service/internal/services/outlet"
	"github.com/mini-e-commerce-microservice/product-service/internal/services/product"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/primitive"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type Dependency struct {
	ProductService product.Service
	OutletService  outlet.Service
}

func NewDependency(appConf *secret_proto.SellerService) (*Dependency, primitive.CloseFn) {
	otelConf := conf.LoadOtelConf()
	minioConf := conf.LoadMinioConf()

	otelCleanup := infra.NewOtel(otelConf, "product-services")
	minioClient := infra.NewMinio(minioConf)
	pgdb, pgdbCleanup := infra.NewPostgresql(appConf.DatabaseDsn)
	rdbms := wsqlx.NewRdbms(pgdb, wsqlx.WithAttributes(semconv.DBSystemPostgreSQL))
	s3Minio := s3_wrapper_minio.New(minioClient)

	// REPO LAYER
	outboxRepository := outbox.New(rdbms)
	productRepository := products.New(rdbms)
	productMediaRepository := product_medias.New(rdbms)
	productVariantRepository := product_variants.New(rdbms)
	productVariantValueRepository := product_variant_values.New(rdbms)
	productVariantItemRepository := product_variant_items.New(rdbms)
	subCategoryItemRepository := sub_category_items.New(rdbms)
	outletRepository := outlets.New(rdbms)

	// SERVICE LAYER
	productService := product.New(product.ServiceOption{
		OutletRepository:              outletRepository,
		SubCategoryItemRepository:     subCategoryItemRepository,
		ProductRepository:             productRepository,
		ProductMediaRepository:        productMediaRepository,
		ProductVariantRepository:      productVariantRepository,
		ProductVariantItemRepository:  productVariantItemRepository,
		ProductVariantValueRepository: productVariantValueRepository,
		OutboxRepository:              outboxRepository,
		S3:                            s3Minio,
		DBTransaction:                 rdbms,
		MinioConf:                     minioConf,
	})
	outletService := outlet.New(outlet.ServiceOption{
		OutletRepository: outletRepository,
		S3:               s3Minio,
		DBTransaction:    rdbms,
		MinioConf:        minioConf,
	})

	dependency := &Dependency{
		OutletService:  outletService,
		ProductService: productService,
	}

	return dependency, func(ctx context.Context) (err error) {
		_ = otelCleanup(ctx)
		_ = pgdbCleanup(ctx)
		return err
	}
}
