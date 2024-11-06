package outlet

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	s3wrapper "github.com/SyaibanAhmadRamadhan/go-s3-wrapper"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/internal/models"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories"
	"github.com/mini-e-commerce-microservice/product-service/internal/repositories/outlets"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/primitive"
	"time"
)

func (s *service) CreateOutlet(ctx context.Context, input CreateOutletInput) (output CreateOutletOutput, err error) {
	outletOutput, err := s.outletRepository.FindOne(ctx, outlets.FindOneInput{
		UserID: null.IntFrom(input.UserID),
	})
	if err != nil {
		if !errors.Is(err, repositories.ErrDataNotFound) {
			return output, collection.Err(err)
		}
	}
	if outletOutput.Data.ID != 0 {
		return output, ErrYouHaveOutlet
	}

	err = s.dbTransaction.DoTxContext(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted},
		func(ctx context.Context, tx wsqlx.Rdbms) (err error) {
			createOutletOutput, err := s.outletRepository.Create(ctx, outlets.CreateInput{
				Tx: tx,
				Data: models.Outlet{
					UserID:      input.UserID,
					Logo:        input.Logo.V.GeneratedFileName,
					Name:        input.Name,
					Slogan:      input.Slogan.ValueOrZero(),
					Description: input.Description.ValueOrZero(),
				},
			})
			if err != nil {
				return collection.Err(err)
			}
			output.ID = createOutletOutput.ID

			if input.Logo.Valid {
				createPresignedUrlOutput, err := s.s3.PresignedUrlUploadObject(ctx, s3wrapper.PresignedUrlUploadObjectInput{
					BucketName: s.minioConf.PrivateBucket,
					Path:       input.Logo.V.GeneratedFileName,
					MimeType:   string(input.Logo.V.MimeType),
					Checksum:   input.Logo.V.ChecksumSHA256,
					Expired:    10 * time.Minute,
				})
				if err != nil {
					return collection.Err(err)
				}

				output.LogoPresignedUrl = null.ValueFrom(primitive.PresignedFileUploadOutput{
					Identifier:      input.Logo.V.Identifier,
					UploadURL:       createPresignedUrlOutput.URL,
					UploadExpiredAt: createPresignedUrlOutput.ExpiredAt,
					MinioFormData:   createPresignedUrlOutput.MinioFormData,
				})
			}

			return err
		},
	)
	if err != nil {
		return output, collection.Err(err)
	}
	return
}

type CreateOutletInput struct {
	UserID      int64
	Logo        null.Value[primitive.PresignedFileUpload]
	Name        string
	Slogan      null.String
	Description null.String
}

type CreateOutletOutput struct {
	ID               int64
	LogoPresignedUrl null.Value[primitive.PresignedFileUploadOutput]
}
