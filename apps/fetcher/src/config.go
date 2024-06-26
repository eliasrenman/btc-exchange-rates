package main

import (
	"log"
	"os"
)

const DATA_URL = "https://blockchain.info/ticker"

var AMQP_CONNECTION string

func init() {
	AMQP_CONNECTION = os.Getenv("AMQP_URL")
	if AMQP_CONNECTION == "" {
		log.Fatal("AMQP_URL environment variable not set")
	}
}
