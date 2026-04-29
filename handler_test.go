package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type stubCEP struct {
	city string
	err  error
}

func (s stubCEP) Lookup(cep string) (string, error) { return s.city, s.err }

type stubWeather struct {
	tempC float64
	err   error
}

func (s stubWeather) GetTemperatureC(loc string) (float64, error) { return s.tempC, s.err }

func TestHandler_Success(t *testing.T) {
	h := &Handler{
		CEPClient:     stubCEP{city: "São Paulo"},
		WeatherClient: stubWeather{tempC: 28.5},
	}
	req := httptest.NewRequest(http.MethodGet, "/weather?cep=01001000", nil)
	rec := httptest.NewRecorder()
	h.Handle(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var got weatherResponseBody
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got.TempC != 28.5 || got.TempF != 83.3 || got.TempK != 301.5 {
		t.Errorf("unexpected response: %+v", got)
	}
}

func TestHandler_InvalidCEP(t *testing.T) {
	cases := []string{"", "123", "abcdefgh", "12345-678", "123456789"}
	for _, cep := range cases {
		h := &Handler{CEPClient: stubCEP{}, WeatherClient: stubWeather{}}
		req := httptest.NewRequest(http.MethodGet, "/weather?cep="+cep, nil)
		rec := httptest.NewRecorder()
		h.Handle(rec, req)

		if rec.Code != http.StatusUnprocessableEntity {
			t.Errorf("cep=%q: status=%d, want 422", cep, rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "invalid zipcode") {
			t.Errorf("cep=%q: body=%q, want 'invalid zipcode'", cep, rec.Body.String())
		}
	}
}

func TestHandler_NotFound(t *testing.T) {
	h := &Handler{
		CEPClient:     stubCEP{err: ErrCEPNotFound},
		WeatherClient: stubWeather{},
	}
	req := httptest.NewRequest(http.MethodGet, "/weather?cep=99999999", nil)
	rec := httptest.NewRecorder()
	h.Handle(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("status=%d, want 404", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "can not find zipcode") {
		t.Errorf("body=%q, want 'can not find zipcode'", rec.Body.String())
	}
}
