package handler

import (
	"errors"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/generated/api"
	"github.com/mini-e-commerce-microservice/product-service/internal/service/product"
	"net/http"
)

func (h *handler) V1ProductPost(w http.ResponseWriter, r *http.Request) {
	userData, ok := h.getUserFromBearerAuth(w, r)
	if !ok {
		return
	}

	req := api.V1ProductPostRequestBody{}
	if !h.httpOtel.BindBodyRequest(w, r, &req) {
		return
	}

	sizeGuideFileUpload, ok := h.bindUploadFileRequestPtr(w, r, req.SizeGuideImage)
	if !ok {
		return
	}

	medias := make([]product.CreateProductInputProductMedia, 0, len(req.Medias))
	for _, media := range req.Medias {
		mediaFileUpload, valid := h.bindUploadFileRequest(w, r, media.Media)
		if !valid {
			return
		}

		medias = append(medias, product.CreateProductInputProductMedia{
			FileUpload: mediaFileUpload,
			IsPrimary:  media.IsPrimaryMedia,
		})
	}

	createProductInput := product.CreateProductInput{
		UserID:            userData.UserId,
		SubCategoryItemID: req.SubCategoryItemId,
		Condition:         req.Condition,
		MinimumPurchase:   req.MinimumPurchase,
		Description:       req.Description,
		Name:              req.Name,
		VariantName1:      null.StringFromPtr(req.VariantName1),
		VariantName2:      null.StringFromPtr(req.VariantName2),
		ProductItems:      make([]product.CreateProductInputProductItem, 0),
		SizeGuide:         null.ValueFromPtr(sizeGuideFileUpload),
		Medias:            medias,
	}

	for _, item := range req.ProductItems {
		productItemImageUpload, ok := h.bindUploadFileRequestPtr(w, r, req.SizeGuideImage)
		if !ok {
			return
		}

		createProductInput.ProductItems = append(createProductInput.ProductItems, product.CreateProductInputProductItem{
			VariantValue1:    null.StringFromPtr(item.VariantValue1),
			VariantValue2:    null.StringFromPtr(item.VariantValue2),
			Price:            item.Price,
			Stock:            item.Stock,
			SKU:              null.StringFromPtr(item.Sku),
			Weight:           item.Weight,
			PackageLength:    item.PackageLength,
			PackageWidth:     item.PackageWidth,
			PackageHeight:    item.PackageHeight,
			IsPrimaryProduct: item.IsPrimaryProduct,
			IsActive:         item.IsActive,
			Image:            null.ValueFromPtr(productItemImageUpload),
		})
	}

	outputCreateProduct, err := h.serv.productService.CreateProduct(r.Context(), createProductInput)
	if err != nil {
		switch {
		case errors.Is(err, product.ErrOnlyChooseOnePrimaryProduct):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrOnlyChooseOnePrimaryProduct.Error())
		case errors.Is(err, product.ErrOnlyChooseOnePrimaryMedia):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrOnlyChooseOnePrimaryMedia.Error())
		case errors.Is(err, product.ErrMustHavePrimaryMedia):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrMustHavePrimaryMedia.Error())
		case errors.Is(err, product.ErrMustHavePrimaryProduct):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrMustHavePrimaryProduct.Error())
		case errors.Is(err, product.ErrInvalidSubCategoryItem):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrInvalidSubCategoryItem.Error())
		case errors.Is(err, product.ErrMustHaveSizeGuide):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrMustHaveSizeGuide.Error())
		case errors.Is(err, product.ErrVariantValue1IsRequired):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrVariantValue1IsRequired.Error())
		case errors.Is(err, product.ErrVariantValue2IsRequired):
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, product.ErrVariantValue2IsRequired.Error())
		default:
			h.httpOtel.Err(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	resp := api.V1ProductPostResponse{
		Id:           outputCreateProduct.ID,
		MediaUploads: make([]api.FileUploadResponse, 0),
	}

	for _, upload := range outputCreateProduct.MediaUploads {
		resp.MediaUploads = append(resp.MediaUploads, h.bindUploadFileResponse(upload))
	}
	for _, upload := range outputCreateProduct.OptionalImageUploads {
		if upload.Valid {
			resp.MediaUploads = append(resp.MediaUploads, h.bindUploadFileResponse(upload.V))
		}
	}

	h.httpOtel.WriteJson(w, r, http.StatusCreated, resp)
}
