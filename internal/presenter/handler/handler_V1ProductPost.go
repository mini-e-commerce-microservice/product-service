package handler

import (
	"github.com/mini-e-commerce-microservice/product-service/generated/api"
	"net/http"
)

func (h *handler) V1ProductPost(w http.ResponseWriter, r *http.Request) {
	req := api.V1ProductPostRequestBody{}
	if !h.httpOtel.BindBodyRequest(w, r, &req) {
		return
	}

}
