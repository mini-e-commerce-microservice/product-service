package handler

import (
	"errors"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/product-service/generated/api"
	"github.com/mini-e-commerce-microservice/product-service/internal/services/outlet"
	"net/http"
)

func (h *handler) V1SellerPost(w http.ResponseWriter, r *http.Request) {
	userData, ok := h.getUserFromBearerAuth(w, r, true)
	if !ok {
		return
	}

	if !userData.IsEmailVerified {
		h.httpOtel.Err(w, r, http.StatusForbidden, errors.New("email user must be verified"), "You must activate your email first")
		return
	}

	req := api.V1SellerPostRequestBody{}
	if !h.httpOtel.BindBodyRequest(w, r, &req) {
		return
	}

	logo, ok := h.bindUploadFileRequestPtr(w, r, req.Logo)
	if !ok {
		return
	}

	createOutletOutput, err := h.serv.outletService.CreateOutlet(r.Context(), outlet.CreateOutletInput{
		UserID:      userData.UserId,
		Name:        req.Name,
		Slogan:      null.StringFromPtr(req.Slogan),
		Description: null.StringFromPtr(req.Description),
		Logo:        null.ValueFromPtr(logo),
	})
	if err != nil {
		if errors.Is(err, outlet.ErrYouHaveOutlet) {
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, outlet.ErrYouHaveOutlet.Error())
		} else {
			h.httpOtel.Err(w, r, http.StatusInternalServerError, err, "Failed to create outlet")
		}
		return
	}

	resp := api.V1SellerPostResponse{
		Id:         createOutletOutput.ID,
		LogoUpload: h.bindUploadFileResponsePtr(createOutletOutput.LogoPresignedUrl.Ptr()),
	}

	h.httpOtel.WriteJson(w, r, http.StatusCreated, resp)
}
