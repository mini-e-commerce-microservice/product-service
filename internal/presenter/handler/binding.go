package handler

import (
	"github.com/mini-e-commerce-microservice/product-service/generated/api"
	"github.com/mini-e-commerce-microservice/product-service/internal/util/primitive"
	"net/http"
)

func (h *handler) bindUploadFileRequest(w http.ResponseWriter, r *http.Request, input api.FileUploadRequest) (output primitive.PresignedFileUpload, ok bool) {
	fileUploadInput := primitive.NewPresignedFileUploadInput{
		Identifier:       input.Identifier,
		OriginalFileName: input.OriginalFilename,
		MimeType:         primitive.MimeType(input.MimeType),
		Size:             input.Size,
		ChecksumSHA256:   input.ChecksumSha256,
	}
	output, err := primitive.NewPresignedFileUpload(fileUploadInput)
	if err != nil {
		h.httpOtel.Err(w, r, http.StatusBadRequest, err, err.Error())
		return output, false
	}
	return output, true
}

func (h *handler) bindUploadFileRequestPtr(w http.ResponseWriter, r *http.Request, input *api.FileUploadRequest) (output *primitive.PresignedFileUpload, ok bool) {
	if input == nil {
		return nil, true
	}

	fileUploadInput := primitive.NewPresignedFileUploadInput{
		Identifier:       input.Identifier,
		OriginalFileName: input.OriginalFilename,
		MimeType:         primitive.MimeType(input.MimeType),
		Size:             input.Size,
		ChecksumSHA256:   input.ChecksumSha256,
	}
	outputNewPresignedFileUpload, err := primitive.NewPresignedFileUpload(fileUploadInput)
	if err != nil {
		h.httpOtel.Err(w, r, http.StatusBadRequest, err, err.Error())
		return output, false
	}
	return &outputNewPresignedFileUpload, true
}

func (h *handler) bindUploadFileResponse(input primitive.PresignedFileUploadOutput) (output api.FileUploadResponse) {
	return api.FileUploadResponse{
		Identifier:      input.Identifier,
		UploadExpiredAt: input.UploadExpiredAt,
		UploadUrl:       input.UploadURL,
		MinioFormData:   input.MinioFormData,
	}
}

func (h *handler) bindUploadFileResponsePtr(input *primitive.PresignedFileUploadOutput) (output *api.FileUploadResponse) {
	if input == nil {
		return nil
	}

	return &api.FileUploadResponse{
		Identifier:      input.Identifier,
		UploadExpiredAt: input.UploadExpiredAt,
		UploadUrl:       input.UploadURL,
		MinioFormData:   input.MinioFormData,
	}
}
