package server

import (
	"currencyApi/currency"
	"currencyApi/server/router"
	"net/http"
)

func CurrencyRoutes(currencyService currency.CurrencyService) []router.Route {
	return []router.Route{
		{
			Path: "/currency",
			Method: http.MethodGet,
			Handler: List(currencyService),
		},
		{
			Path: "/currency/:code",
			Method: http.MethodGet,
			Handler: GetByCode(currencyService),
		},
		{
			Path: "/currency",
			Method: http.MethodPost,
			Handler: Create(currencyService),
		},
	}
}
