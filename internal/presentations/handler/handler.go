package handler

import (
	whttp "github.com/SyaibanAhmadRamadhan/http-wrapper"
	"github.com/go-chi/chi/v5"
	"github.com/mini-e-commerce-microservice/product-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/product-service/internal/services/outlet"
	"github.com/mini-e-commerce-microservice/product-service/internal/services/product"
)

type handler struct {
	r                  *chi.Mux
	httpOtel           *whttp.Opentelemetry
	serv               serv
	jwtAccessTokenConf *secret_proto.JwtAccessToken
}

type serv struct {
	productService product.Service
	outletService  outlet.Service
}

type Opt struct {
	ProductService     product.Service
	OutletService      outlet.Service
	JwtAccessTokenConf *secret_proto.JwtAccessToken
}

func Init(r *chi.Mux, opt Opt) {
	h := &handler{
		r: r,
		httpOtel: whttp.NewOtel(
			whttp.WithRecoverMode(true),
			whttp.WithPropagator(),
			whttp.WithValidator(nil, nil),
		),
		jwtAccessTokenConf: opt.JwtAccessTokenConf,
		serv: serv{
			productService: opt.ProductService,
			outletService:  opt.OutletService,
		},
	}
	h.route()
}

func (h *handler) route() {
	h.r.Post("/v1/product", h.httpOtel.Trace(
		h.V1ProductPost,
	))

	h.r.Post("/v1/seller", h.httpOtel.Trace(
		h.V1SellerPost,
	))
}
