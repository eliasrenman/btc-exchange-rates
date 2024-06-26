package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type ConversionRates = map[string]ConversionRate

type ConversionRate struct {
	FifteenMinutes float64 `json:"15m"`
	Last           float64 `json:"last"`
	Buy            float64 `json:"buy"`
	Sell           float64 `json:"sell"`
	Symbol         string  `json:"symbol"`
}

func fetchLatestExchangeRate() (ConversionRates, error) {
	resp, err := http.Get(DATA_URL)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result ConversionRates
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalln("Failed unwrapping response body, error:", err)
		return result, err
	}
	return result, nil
}
