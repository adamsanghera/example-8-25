package web

import (
	"context"
	"net/http"
)

type Web struct {
	srv *http.Server
}

func New(cfg *Config) (*Web, error) {
	return &Web{
		srv: &http.Server{
			Addr: cfg.Addr,
		},
	}, nil
}

func (w *Web) Start() error {
	mux := http.NewServeMux()
	w.srv.Handler = mux

	mux.HandleFunc("/", index)

	return w.srv.ListenAndServe()
}

func (w *Web) Shutdown(ctx context.Context) error {
	return w.srv.Shutdown(ctx)
}
