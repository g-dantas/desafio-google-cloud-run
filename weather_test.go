package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWeatherAPIClient_GetTemperatureC(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("key") != "test-key" {
			t.Errorf("missing api key")
		}
		if r.URL.Query().Get("q") != "São Paulo" {
			t.Errorf("unexpected q = %q", r.URL.Query().Get("q"))
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"current":{"temp_c":22.5}}`))
	}))
	defer server.Close()

	c := &WeatherAPIClient{HTTPClient: server.Client(), BaseURL: server.URL, APIKey: "test-key"}
	temp, err := c.GetTemperatureC("São Paulo")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if temp != 22.5 {
		t.Errorf("temp = %v, want 22.5", temp)
	}
}
