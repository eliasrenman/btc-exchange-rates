package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type ExchangeRateMessage struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func insertRow(message ExchangeRateMessage) {
	conn, err := pgx.Connect(context.Background(), POSTGRES_CONNECTION)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	query := `INSERT INTO exchange_rates (currency, rate) VALUES ($1, $2)`

	// Execute the query
	_, err = conn.Exec(context.Background(), query, message.Currency, message.Rate)

	if err != nil {
		log.Fatalf("Failed to insert row: %v\n", err)
		os.Exit(1)
	}
	log.Println("Successfully inserted latest exchange rate")
}
