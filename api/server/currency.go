package server

import (
	"currencyApi/currency"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Currency struct {
	Name string  `json:"name,omitempty"`
	Code string  `json:"code"`
	Rate float32 `json:"rate"`
}

type CurrencyListResponse struct {
	Data []Currency
}

type CurrencyCreationRequest struct {
	Name string  								`json:"name"`
	Code string  								`json:"code"`
	Rate float32 								`json:"rate"`
	Type currency.CurrencyType  `json:"type"`
}

func List(currencyService currency.CurrencyService) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		currencies, err := currencyService.Find()
		if err != nil {
			panic(err)
		}

		data := []Currency{}
		for _, c := range currencies {
			data = append(data, Currency{
				Name: c.Name,
				Code: c.Code,
				Rate: c.Rate,
			})
		}

		response := CurrencyListResponse{data}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})
}

func GetByCode(currencyService currency.CurrencyService) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		codeParam := params.ByName("code")
		currency, err := currencyService.FindOne(codeParam)
		if err != nil {
			panic(err)
		}

		if currency == nil {
			panic(errors.New("not found."))
		}

		if err := json.NewEncoder(w).Encode(currency); err != nil {
			panic(err)
		}
	})
}

func Create(currencyService currency.CurrencyService) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		var requestData CurrencyCreationRequest
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			panic(err)
		}

		data := currencyFromRequest(requestData)
		currency, err := currencyService.Create(data)
		if err != nil {
			panic(err)
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(currency); err != nil {
			panic(err)
		}
	})
}

func currencyFromRequest(requestData CurrencyCreationRequest) currency.Currency {
	return currency.Currency{
		Name: requestData.Name,
		Code: requestData.Code,
		Rate: requestData.Rate,
		Type: requestData.Type,
	}
}