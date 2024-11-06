package outlet

import (
	s3wrapper "github.com/SyaibanAhmadRamadhan/go-s3-wrapper"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/product-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/outlets"
)

type service struct {
	outletRepository outlets.Repository
	s3               s3wrapper.S3Client
	dbTransaction    wsqlx.Tx

	// config
	minioConf *secret_proto.Minio
}

type ServiceOption struct {
	OutletRepository outlets.Repository
	S3               s3wrapper.S3Client
	DBTransaction    wsqlx.Tx

	// config
	MinioConf *secret_proto.Minio
}

func New(opt ServiceOption) *service {
	return &service{
		outletRepository: opt.OutletRepository,
		s3:               opt.S3,
		dbTransaction:    opt.DBTransaction,
		minioConf:        opt.MinioConf,
	}
}
