package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type CEPLookup interface {
	Lookup(cep string) (string, error)
}

type WeatherLookup interface {
	GetTemperatureC(location string) (float64, error)
}

type Handler struct {
	CEPClient     CEPLookup
	WeatherClient WeatherLookup
}

type weatherResponseBody struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if !IsValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := h.CEPClient.Lookup(cep)
	if err != nil {
		if errors.Is(err, ErrCEPNotFound) {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempC, err := h.WeatherClient.GetTemperatureC(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body := weatherResponseBody{
		TempC: tempC,
		TempF: CelsiusToFahrenheit(tempC),
		TempK: CelsiusToKelvin(tempC),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
