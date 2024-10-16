package handler

import (
	whttp "github.com/SyaibanAhmadRamadhan/http-wrapper"
	"github.com/go-chi/chi/v5"
	"github.com/mini-e-commerce-microservice/product-service/internal/service/product"
)

type handler struct {
	r        *chi.Mux
	httpOtel *whttp.Opentelemetry
	serv     serv
}

type serv struct {
	productService product.Service
}

type Opt struct {
	ProductService product.Service
}

func Init(r *chi.Mux, opt Opt) {
	h := &handler{
		r: r,
		httpOtel: whttp.NewOtel(
			whttp.WithRecoverMode(true),
			whttp.WithPropagator(),
			whttp.WithValidator(nil, nil),
		),
		serv: serv{
			productService: opt.ProductService,
		},
	}
	h.route()
}

func (h *handler) route() {
	h.r.Post("/v1/product", h.httpOtel.Trace(
		h.V1ProductPost,
	))
}
