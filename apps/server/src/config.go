package main

import (
	"log"
	"os"
)

var POSTGRES_CONNECTION string

func init() {
	POSTGRES_CONNECTION = os.Getenv("DATABASE_URL")
	if POSTGRES_CONNECTION == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
}
