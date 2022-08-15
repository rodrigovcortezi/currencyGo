package server

import (
	"context"
	"currencyApi/currency"
	"currencyApi/server/router"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string, currencyService currency.CurrencyService) *Server {
	r := router.New(CurrencyRoutes(currencyService))

	return &Server{
		httpServer: &http.Server{
			Addr: addr,
			Handler: r,
		},
	}
}

func (s Server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
