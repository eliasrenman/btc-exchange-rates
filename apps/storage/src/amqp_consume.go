package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rabbitmq/amqp091-go"
)

func startConsumer() {
	log.Println("Connecting to AMQP")

	conn, err := amqp091.Dial(AMQP_CONNECTION)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	log.Println("Successfully connected to AMQP")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"exchange-rate-queue", // queue name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		queue.Name, // queue name
		"",         // consumer tag
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Handle incoming messages in a separate goroutine
	go func() {
		for msg := range msgs {
			fmt.Printf("Received a message: %s\n", msg.Body)
			// Process the exchange rate update here
			var result ExchangeRateMessage
			if err := json.Unmarshal(msg.Body, &result); err != nil {
				log.Fatalln("Failed unwrapping message body, error:", err)
				continue
			}
			insertRow(result)
		}
	}()

	// Wait for termination signal to gracefully close the connection
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("Shutting down...")
}
