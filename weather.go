package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherAPIClient struct {
	HTTPClient *http.Client
	APIKey     string
	BaseURL    string
}

type weatherResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (c *WeatherAPIClient) GetTemperatureC(location string) (float64, error) {
	baseURL := c.BaseURL
	if baseURL == "" {
		baseURL = "http://api.weatherapi.com/v1/current.json"
	}
	q := url.Values{}
	q.Set("key", c.APIKey)
	q.Set("q", location)

	resp, err := c.HTTPClient.Get(baseURL + "?" + q.Encode())
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("weatherapi returned status %d", resp.StatusCode)
	}

	var data weatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}
	return data.Current.TempC, nil
}
