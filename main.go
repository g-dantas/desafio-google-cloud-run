package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h := &Handler{
		CEPClient:     &ViaCEPClient{HTTPClient: http.DefaultClient},
		WeatherClient: &WeatherAPIClient{HTTPClient: http.DefaultClient, APIKey: apiKey},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/weather", h.Handle)

	log.Printf("server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
