package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValidCEP(t *testing.T) {
	tests := []struct {
		cep  string
		want bool
	}{
		{"12345678", true},
		{"00000000", true},
		{"1234567", false},
		{"123456789", false},
		{"1234567a", false},
		{"", false},
		{"12345-678", false},
	}
	for _, tt := range tests {
		if got := IsValidCEP(tt.cep); got != tt.want {
			t.Errorf("IsValidCEP(%q) = %v, want %v", tt.cep, got, tt.want)
		}
	}
}

func TestViaCEPClient_LookupSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/01001000/json/" {
			t.Errorf("unexpected path %q", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"localidade":"São Paulo"}`))
	}))
	defer server.Close()

	c := &ViaCEPClient{HTTPClient: server.Client(), BaseURL: server.URL}
	city, err := c.Lookup("01001000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if city != "São Paulo" {
		t.Errorf("city = %q, want São Paulo", city)
	}
}

func TestViaCEPClient_LookupNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"erro":true}`))
	}))
	defer server.Close()

	c := &ViaCEPClient{HTTPClient: server.Client(), BaseURL: server.URL}
	_, err := c.Lookup("99999999")
	if !errors.Is(err, ErrCEPNotFound) {
		t.Errorf("err = %v, want ErrCEPNotFound", err)
	}
}
