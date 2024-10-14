package handler

import (
	whttp "github.com/SyaibanAhmadRamadhan/http-wrapper"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	r        *chi.Mux
	httpOtel *whttp.Opentelemetry
}

type Opt struct{}

func Init(r *chi.Mux, opt Opt) {
	h := &handler{
		r: r,
		httpOtel: whttp.NewOtel(
			whttp.WithRecoverMode(true),
			whttp.WithPropagator(),
			whttp.WithValidator(nil, nil),
		),
	}
	h.route()
}

func (h *handler) route() {
	h.r.Post("/v1/product", h.httpOtel.Trace(
		h.V1ProductPost,
	))
}
