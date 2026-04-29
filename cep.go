package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

var ErrCEPNotFound = errors.New("cep not found")

var cepRegex = regexp.MustCompile(`^\d{8}$`)

func IsValidCEP(cep string) bool {
	return cepRegex.MatchString(cep)
}

type ViaCEPClient struct {
	HTTPClient *http.Client
	BaseURL    string
}

type viaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       any    `json:"erro,omitempty"`
}

func (c *ViaCEPClient) Lookup(cep string) (string, error) {
	baseURL := c.BaseURL
	if baseURL == "" {
		baseURL = "https://viacep.com.br/ws"
	}
	url := fmt.Sprintf("%s/%s/json/", baseURL, cep)

	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrCEPNotFound
	}

	var data viaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}
	if data.Erro != nil || data.Localidade == "" {
		return "", ErrCEPNotFound
	}
	return data.Localidade, nil
}
