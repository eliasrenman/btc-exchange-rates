package main

import (
	"log"
	"os"
)


var AMQP_CONNECTION string

func init() {
	AMQP_CONNECTION = os.Getenv("AMQP_URL")
	if AMQP_CONNECTION == "" {
		log.Fatal("AMQP_URL environment variable not set")
	}
}
