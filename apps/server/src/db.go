package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

// ExchangeRate represents a single row from the exchange_rates table
type ExchangeRate struct {
	Currency  string    `json:"currency"`
	Rate      float64   `json:"rate"`
	CreatedAt time.Time `json:"createdAt"`
}

func queryLatestExchange() (*ExchangeRate, error) {

	// Create a new PostgreSQL connection
	conn, err := pgx.Connect(context.Background(), POSTGRES_CONNECTION)
	if err != nil {
		log.Fatalln("unable to connect to database", err)
	}
	defer conn.Close(context.Background())
	query := "SELECT currency, rate, created_at FROM exchange_rates ORDER BY created_at DESC LIMIT 1"
	// Query the exchange_rates table for the latest entry
	var exchangeRate ExchangeRate
	err = conn.QueryRow(context.Background(), query).
		Scan(&exchangeRate.Currency, &exchangeRate.Rate, &exchangeRate.CreatedAt)
	if err != nil {
		log.Println("query failed", err)
	}

	return &exchangeRate, nil
}
