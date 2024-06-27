package main

import (
	"log"
	"os"
)

var AMQP_CONNECTION string
var POSTGRES_CONNECTION string

func init() {
	AMQP_CONNECTION = os.Getenv("AMQP_URL")
	if AMQP_CONNECTION == "" {
		log.Fatal("AMQP_URL environment variable not set")
	}

	POSTGRES_CONNECTION = os.Getenv("DATABASE_URL")
	if POSTGRES_CONNECTION == "" {
		log.Fatal("POSTGRES_CONNECTION environment variable not set")
	}
}
